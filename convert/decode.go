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
func Decode(w io.Writer, r io.Reader) error {

	// read char by char because of "not load the whole file into memory at once"
	// reading tokens wouldn't work for huge files consisting only of non-space
	mr := morseReader{br: bufio.NewReader(r)}

	for {
		m, err := mr.readMorse()

		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return err
		}

		toWrite := ""
		if m == "/" {
			toWrite = " "
		} else if m == "//" {
			toWrite = "\n"
		} else if toWrite, err = fromMorse(m); err != nil {
			return err
		}

		if _, err = io.WriteString(w, toWrite); err != nil {
			return err
		}
	}
}

type morseReader struct {
	br   *bufio.Reader
	next string // holds incomplete morse sequence
}

// readMorse tries to read valid morse. If successful, it returns
// a sequence of '.' and '-' or "/" or "//"
func (mr *morseReader) readMorse() (string, error) {
	for {
		r, _, err := mr.br.ReadRune()

		if err != nil {
			// return EOF only if mr.next is not empty
			// otherwise next call to readMorse will return EOF
			if mr.next == "" {
				return "", err
			}
			morse := mr.next
			mr.next = ""
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return morse, err
		}

		switch {
		case r == '.' || r == '-': // inside morse sequence
			if mr.next == "/" {
				morse := mr.next
				mr.next = string(r)
				return morse, nil
			}
			mr.next += string(r)

		case unicode.IsSpace(r): // character separator
			morse := mr.next
			mr.next = ""
			if morse != "" {
				return morse, nil
			}

		case r == '/': // whitespace
			if mr.next == "/" {
				morse := "//"
				mr.next = ""
				return morse, nil
			}
			morse := mr.next
			mr.next = "/"
			return morse, nil

		default:
			return mr.next, ErrInvalidMorse
		}
	}
}
