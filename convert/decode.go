package convert

import (
	"errors"
	"io"
)

// Decode creates utf8 from morse.
// It fails on any non EOF error when reading from w.
// And if w contains invalid UTF8.
// Characters for which no morse code exists are silently dropped.
func Decode(w io.Writer, r io.Reader) error {
	return errors.New("not implemented")
}
