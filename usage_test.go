package input

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func printEvent(e *Event) {
	log.Printf("got event: %v", e)
}

func Example() {
	// create processor
	ip := NewProcessor()

	// register specific event
	motion := TypeCode{
		Type: EV_KEY,
		Code: KEY_SPACE,
	}
	ip.Register(motion, printEvent)

	// register all events of a type
	dial := TypeCode{
		Type: EV_REL,
		Code: REL_CNT,
	}
	ip.Register(dial, printEvent)

	// open event devices
	fns, _ := filepath.Glob("/dev/input/event*")
	var fs []io.ReadCloser
	for _, fn := range fns {
		f, _ := os.Open(fn)
		fs = append(fs, f)
	}

	// Run Watch and Process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = ip.WatchAndProcess(ctx, fs)
}
