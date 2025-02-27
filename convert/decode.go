package convert

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

var ErrInvalidMorse = errors.New("invalid morse")

// Decode creates utf8 from morse.
// It fails on any non EOF error when reading from w.
// And if w contains invalid UTF8.
// Characters for which no morse code exists are silently dropped.
func Decode(w io.Writer, r io.Reader) (err error) {

	// read char by char because of "not load the whole file into memory at once"
	// reading tokens wouldn't work for huge files consisting only of non-space
	br := bufio.NewReader(r)
	currentMorse := "" // todo: or use bytes.Buffer?
	var lastChar rune  // to detect 2nd slash

	for err == nil {
		var r rune
		toWrite := ""

		if r, _, err = br.ReadRune(); err != nil {
			break // EOF or other error - checked outside loop

		} else if r == '.' || r == '-' { // inside a morse code
			if lastChar == '/' {
				toWrite = " "
				lastChar = 0
			}
			currentMorse += string(r)

		} else if unicode.IsSpace(r) { // after a morse code
			if toWrite, err = fromMorse(currentMorse); err != nil {
				return err
			}
			if lastChar == '/' { // or is space after slash invalid?
				toWrite = " " + toWrite
			}
			currentMorse = ""
			lastChar = 0

		} else if r == '/' { // word or line separator
			if lastChar == '/' {
				toWrite = "\n"
				lastChar = 0
			} else {
				lastChar = r
			}
		} else {
			return ErrInvalidMorse
		}

		if toWrite != "" {
			if _, err = io.WriteString(w, toWrite); err != nil {
				return err
			}
		}
	}

	if !errors.Is(err, io.EOF) {
		return err
	}

	if lastChar == '/' {
		if _, err = io.WriteString(w, " "); err != nil {
			return err
		}
	}

	toWrite := ""
	if toWrite, err = fromMorse(currentMorse); err == nil {
		_, err = io.WriteString(w, toWrite)
	}
	return err
}
