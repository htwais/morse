// Package convert provides functions to create morse from utf8 and vice versa.
package convert

import (
	"errors"
	"io"
)

// Encode creates morse from utf8 text.
// It fails on any non EOF error when reading from w.
// And if w contais characters outside the valid range - [0x20, 0x2d, 0x2e, 0x2f].
func Encode(w io.Writer, r io.Reader) error {
	return errors.New("not implemented")
}
