package input

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

//go:generate stringer -type=EV
type EV uint16

type Coder interface {
	Code() CODE
	String() string
}

type CODE uint16

func (c CODE) Code() CODE {
	return c
}

func (c CODE) String() string {
	return strconv.Itoa(int(c))
}

//go:generate stringer -type=REL_CODE
type REL_CODE CODE

func (c REL_CODE) Code() CODE {
	return CODE(c)
}

//go:generate stringer -type=KEY_CODE
type KEY_CODE CODE

func (c KEY_CODE) Code() CODE {
	return CODE(c)
}

type TypeCode struct {
	Type EV
	Code Coder
}

// Event is a linux subsystem event
type Event struct {
	Time  time.Time
	Type  EV
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

	ev := EV(binary.LittleEndian.Uint16(bs[8:10]))
	if ev >= EV_CNT {
		return nil, ErrDecode{fmt.Sprintf("invalid event: %v", ev)}
	}

	sec := binary.LittleEndian.Uint32(bs[0:4])
	usec := binary.LittleEndian.Uint32(bs[4:8])

	c := binary.LittleEndian.Uint16(bs[10:12])

	var code Coder
	switch ev {
	case EV_REL:
		code = REL_CODE(c)
	case EV_KEY:
		code = KEY_CODE(c)
	default:
		code = CODE(c)
	}

	e := &Event{
		Time:  time.Unix(int64(sec), int64(usec)),
		Type:  ev,
		Code:  code,
		Value: int32(binary.LittleEndian.Uint32(bs[12:16])),
	}

	return e, nil
}

func (e *Event) String() string {
	return fmt.Sprintf("%s %s %s %d", e.Time.UTC().Format(time.RFC3339), e.Type, e.Code, e.Value)
}
