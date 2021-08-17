package input

import "fmt"

type ErrDecode struct {
	s string
}

func (e ErrDecode) Error() string {
	return fmt.Sprintf("decode error: %v", e.s)
}
