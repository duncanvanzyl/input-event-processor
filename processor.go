package input

import (
	"context"
	"io"
	"sync"
)

// type Handler interface {
// 	Handle(*Event)
// }

// HandleFunc is a callback that does something on receipt of an event when
// registered with a Processor using the Register function.
//
// HandleFunc needs to check event value as required.
type HandleFunc func(*Event)

// func (hf HandleFunc) Handle(e *Event) {
// 	hf(e)
// }

// Processor can watch a Linux Event file for input and call a HandleFunc on
// matched input.
//
// Register hanlders with Register.
//
// Run with WatchAndProcess.
type Processor struct {
	events map[TypeCode][]HandleFunc
}

// NewProcessor returns an initialized processor
func NewProcessor() *Processor {
	return &Processor{events: make(map[TypeCode][]HandleFunc)}
}

// Register registers a HandlerFunc to be called when an event with the given TypeCode arrives.
//
// Specific events, such as individual key presses, are called by providing both the Type and Code.
// eg:
//   Type   Code
//   EV_KEY KEY_SPACE
//
// All events of a type can be handled by a single HandleFunc by setting the code to the _CNT value
// for that type.
// eg:
//   Type   Code
//   EV_KEY KEY_CNT
//
// Event value is not considered when processing event. HandlerFunc should check the event value as
// required.
func (p *Processor) Register(tc TypeCode, f HandleFunc) {
	p.events[tc] = append(p.events[tc], f)
}

// WatchAndProcess Reads events from the readers, and calls any functions registered to the event.
// Blocks until an error occurs, all the readers close, or until the context is canceled. Calls a
// seperate go routine to monitor each io.Reader provided. Calls registered functions in their own
// go routines.
// Register events with processor's Register function.
func (p *Processor) WatchAndProcess(ctx context.Context, rs []io.ReadCloser) error {
	// logger := hclog.FromContext(ctx).Named("input")

	type eventError struct {
		e   *Event
		err error
	}

	c := make(chan eventError, 1)
	wg := sync.WaitGroup{}

	// read from all readers in seperate go routines until either the reader closes
	// if there is an error, push the error onto the channel
	for _, r := range rs {
		wg.Add(1)
		go func(r io.Reader) {
			defer wg.Done()
			for {
				b := make([]byte, 16)
				_, err := r.Read(b)
				if err != nil {
					if err != io.EOF {
						c <- eventError{nil, err}
					}
					return
				}

				e, err := FromBytes(b)
				if err != nil {
					c <- eventError{nil, err}
				}
				c <- eventError{e, nil}
			}
		}(r)
	}

	// wait for all the readers to close, and then close the channel
	go func() {
		wg.Wait()
		close(c)
	}()

	closeReaders := func() {
		for _, r := range rs {
			r.Close()
		}
	}

	// if context is done, then close the readers
	go func(ctx context.Context) {
		<-ctx.Done()
		closeReaders()
	}(ctx)

	// read from the channel. if an error arrives, then close the readers, which will in turn close
	// the channel, which will end this loop.
	var err error
	for ee := range c {
		if ee.err != nil {
			err = ee.err
			closeReaders()
		}

		if e := ee.e; e != nil {
			// handle specific code (ie, single key press)
			tc := TypeCode{e.Type, e.Code}
			if fs, ok := p.events[tc]; ok {
				for _, f := range fs {
					go f(e)
				}
			}

			// handle all code for an event type
			switch tc.Type {
			case EV_REL:
				tc.Code = REL_CNT
			case EV_KEY:
				tc.Code = KEY_CNT
			}
			if fs, ok := p.events[tc]; ok {
				for _, f := range fs {
					go f(e)
				}
			}
		}
	}

	return err
}
