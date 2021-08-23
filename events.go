package input

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

type Coder interface {
	Code() Code
	String() string
}

type Code uint16

func (c Code) Code() Code {
	return c
}

func (c Code) String() string {
	return strconv.Itoa(int(c))
}

//go:generate go run generator.go https://raw.githubusercontent.com/torvalds/linux/master/include/uapi/linux/input-event-codes.h codes.go

type TypeCode struct {
	Type EVCode
	Code Coder
}

// Event is a linux subsystem event
type Event struct {
	Time  time.Time
	Type  EVCode
	Code  Coder
	Value int32
}

// FromBytes returns an event and and error when trying to decode an event from bytes.
// Note that this is designed for a 32 bit system (Raspberry Pi), so the time value, and all other
// bytes are based on a 16 byte slice.
func FromBytes(bs []byte) (*Event, error) {
	// TODO:
	//  - consider 64 bit operating systems with 24 byte slices.
	//  - is this the most efficient way to do this? Maybe just decode directly into a struct
	//  - maybe check for input system requirements (ie, check that it is running on linux)
	if len(bs) != 16 {
		return nil, ErrDecode{fmt.Sprintf("invalid event: should be %d bytes is %d bytes", 16, len(bs))}
	}

	ev := EVCode(binary.LittleEndian.Uint16(bs[8:10]))
	if ev >= EV_CNT {
		return nil, ErrDecode{fmt.Sprintf("invalid event: %v", ev)}
	}

	sec := binary.LittleEndian.Uint32(bs[0:4])
	usec := binary.LittleEndian.Uint32(bs[4:8])

	c := binary.LittleEndian.Uint16(bs[10:12])

	return &Event{
		Time:  time.Unix(int64(sec), int64(usec)),
		Type:  ev,
		Code:  evCode(ev, c),
		Value: int32(binary.LittleEndian.Uint32(bs[12:16])),
	}, nil
}

func (e *Event) String() string {
	return fmt.Sprintf("%s %s %s %d", e.Time.UTC().Format(time.RFC3339), e.Type, e.Code, e.Value)
}
