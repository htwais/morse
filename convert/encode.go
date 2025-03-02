// Package convert provides functions to create morse from utf8 and vice versa.
package convert

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

var ErrInvalidUTF8 = errors.New("invalid UTF8")

const wordSeparator = "/"
const lineSeparator = "//"

// Encode creates morse from utf8 text.
// It fails on any non EOF error when reading from w.
// And if w contais characters outside the valid range - [0x20, 0x2d, 0x2e, 0x2f].
func Encode(w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)
	letterSeparator := "" // only very first morse is without space prefix

	for {
		r, _, err := br.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		// ReadRune returns ReplacementChar for invalid UTF8:
		if r == unicode.ReplacementChar {
			return ErrInvalidUTF8
		}

		// handle whitespace: output '/' or '//'
		// todo: \r\n line endings will output space and // - is that ok?
		toWrite := ""
		if r == '\n' {
			toWrite = letterSeparator + lineSeparator
		} else if unicode.IsSpace(r) {
			toWrite = letterSeparator + wordSeparator
		} else {
			toWrite = letterSeparator + toMorse(r)
		}
		letterSeparator = " "

		if _, err = io.WriteString(w, toWrite); err != nil {
			return err
		}
	}
}
