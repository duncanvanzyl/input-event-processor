package input

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"
)

// TODO: Add tests
//  - closing reader twice
//  - cancel context to exit
//  - multiple readers
//  - error on read
//  - error on decode

func TestProcessor(t *testing.T) {

	tt := []struct {
		name          string
		tc            TypeCode
		bs            []byte
		dontClose     bool
		expectedValue string
		expectedError error
	}{
		{"Space Key Down",
			TypeCode{EV_KEY, KEY_SPACE},
			[]byte{209, 200, 229, 95, 0, 0, 0, 0, // time
				1, 0, // EV_KEY
				57, 0, // KEY_SPACE
				1, 0, 0, 0}, // Key down
			false,
			"2020-12-25T11:11:13Z EV_KEY KEY_SPACE 1", nil},
		{"Space Key Down - All Keys",
			TypeCode{EV_KEY, KEY_CNT},
			[]byte{209, 200, 229, 95, 0, 0, 0, 0, // time
				1, 0, // EV_KEY
				57, 0, // KEY_SPACE
				1, 0, 0, 0}, // Key down
			false,
			"2020-12-25T11:11:13Z EV_KEY KEY_SPACE 1", nil},
		{"Space Key Down - Not Listening",
			TypeCode{EV_REL, REL_DIAL},
			[]byte{209, 200, 229, 95, 0, 0, 0, 0, // time
				1, 0, // EV_KEY
				57, 0, // KEY_SPACE
				1, 0, 0, 0}, // Key down
			false,
			"", nil},
		{"Dial Up",
			TypeCode{EV_REL, REL_DIAL},
			[]byte{209, 200, 229, 95, 0, 0, 0, 0, // time
				2, 0, // EV_REL
				7, 0, // REL_DIAL
				1, 0, 0, 0}, // Dial Clockwise
			false,
			"2020-12-25T11:11:13Z EV_REL REL_DIAL 1", nil},
		{"Dial Down",
			TypeCode{EV_REL, REL_DIAL},
			[]byte{209, 200, 229, 95, 0, 0, 0, 0, // time
				2, 0, // EV_REL
				7, 0, // REL_DIAL
				255, 255, 255, 255}, // Dial Clockwise
			false,
			"2020-12-25T11:11:13Z EV_REL REL_DIAL -1", nil},
		{"Wrong Length", TypeCode{EV_KEY, KEY_SPACE}, []byte{0x00}, false, "", nil},
		{"Invalid Event Code", TypeCode{EV_KEY, KEY_SPACE}, []byte{209, 200, 229, 95, 0, 0, 0, 0, // time
			0x20, 0, // EV_CNT
			7, 0, // REL_DIAL
			255, 255, 255, 255}, // Dial Clockwise
			false,
			"", ErrDecode{}},
		{"Pipe Error", TypeCode{EV_KEY, KEY_SPACE}, []byte{0x00}, true, "", io.ErrClosedPipe},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			p := NewProcessor()

			s := ""

			p.Register(test.tc, func(e *Event) { s += e.String() })

			r, w := io.Pipe()

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
			defer cancel()

			go func(t *testing.T) {
				_, err := w.Write(test.bs)
				if err != nil {
					t.Errorf("could not write to buffer: %v", err)
				}
				if !test.dontClose {
					w.Close()
				}
			}(t)

			err := p.WatchAndProcess(ctx, []io.ReadCloser{r})
			if err != nil {
				if fmt.Sprintf("%T", err) != fmt.Sprintf("%T", test.expectedError) {
					t.Log(fmt.Sprintf("%T-%T", err, test.expectedError))
					t.Fatalf("watch and process failed: %v, %T", err, err)
				}
			}

			<-ctx.Done()

			if s != test.expectedValue {
				t.Errorf("incorrect value: expected %q, got %q", test.expectedValue, s)
			}
		})
	}
}
