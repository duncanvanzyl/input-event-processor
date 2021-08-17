package input

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	et := time.Date(2020, 12, 25, 11, 11, 13, 10, time.UTC)

	tt := []struct {
		name          string
		event         *Event
		expectedValue string
	}{
		{"Space Key Down", &Event{et, EV_KEY, KEY_SPACE, 1}, "2020-12-25T11:11:13Z EV_KEY KEY_SPACE 1"},
		{"Dial Up", &Event{et, EV_REL, REL_DIAL, 1}, "2020-12-25T11:11:13Z EV_REL REL_DIAL 1"},
		{"Dial Down", &Event{et, EV_REL, REL_DIAL, -1}, "2020-12-25T11:11:13Z EV_REL REL_DIAL -1"},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			s := test.event.String()
			if s != test.expectedValue {
				t.Errorf("string failed: expected %q, got %q", test.expectedValue, s)
			}
		})
	}
}

func TestFromBytes(t *testing.T) {
	// Atleast for my purposes:
	// time.Date(2020, 12, 25, 11, 11, 13, 10, time.UTC) == []byte{209, 200, 229, 95, 0, 0, 0, 0}

	tt := []struct {
		name          string
		bs            []byte
		expectedValue string
		shouldFail    bool
	}{
		{"Space Key Down", []byte{209, 200, 229, 95, 0, 0, 0, 0, // time
			1, 0, // EV_KEY
			57, 0, // KEY_SPACE
			1, 0, 0, 0}, // Key down
			"2020-12-25T11:11:13Z EV_KEY KEY_SPACE 1", false},
		{"Dial Up", []byte{209, 200, 229, 95, 0, 0, 0, 0, // time
			2, 0, // EV_REL
			7, 0, // REL_DIAL
			1, 0, 0, 0}, // Dial Clockwise
			"2020-12-25T11:11:13Z EV_REL REL_DIAL 1", false},
		{"Dial Down", []byte{209, 200, 229, 95, 0, 0, 0, 0, // time
			2, 0, // EV_REL
			7, 0, // REL_DIAL
			255, 255, 255, 255}, // Dial Clockwise
			"2020-12-25T11:11:13Z EV_REL REL_DIAL -1", false},
		{"Wrong Length", []byte{}, "", true},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			e, err := FromBytes(test.bs)
			if (err != nil) != test.shouldFail {
				if test.shouldFail {
					t.Fatalf("should have failed but did not: %v", err)
				} else {
					t.Fatalf("should not have failed but did: %v", err)
				}
			}

			if !test.shouldFail {
				s := e.String()
				if s != test.expectedValue {
					t.Errorf("string failed: expected %q, got %q", test.expectedValue, s)
				}
			}
		})
	}
}
