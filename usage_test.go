package input_test

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/duncanvanzyl/input-event-processor"
)

func printEvent(e *input.Event) {
	log.Printf("got event: %v", e)
}

func Example() {
	// create processor
	ip := input.NewProcessor()

	// register specific event
	motion := input.TypeCode{
		Type: input.EV_KEY,
		Code: input.KEY_SPACE,
	}
	ip.Register(motion, printEvent)

	// register all events of a type
	dial := input.TypeCode{
		Type: input.EV_REL,
		Code: input.REL_CNT,
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
