package convert

import (
	"errors"
	"unicode"
)

var ErrUnknownMorse = errors.New("unknown morse")

// a hardcoded table of (most?) morse codes and their text mappings:
var m2t map[string]rune = map[string]rune{
	".-":     'a',
	"-...":   'b',
	"-.-.":   'c',
	"-..":    'd',
	".":      'e',
	"..-.":   'f',
	"--.":    'g',
	"....":   'h',
	"..":     'i',
	".---":   'j',
	"-.-":    'k',
	".-..":   'l',
	"--":     'm',
	"-.":     'n',
	"---":    'o',
	".--.":   'p',
	"--.-":   'q',
	".-.":    'r',
	"...":    's',
	"-":      't',
	"..-":    'u',
	"...-":   'v',
	".--":    'w',
	"-..-":   'x',
	"-.--":   'y',
	"--..":   'z',
	".----":  '1',
	"..---":  '2',
	"...--":  '3',
	"....-":  '4',
	".....":  '5',
	"-....":  '6',
	"--...":  '7',
	"---..":  '8',
	"----.":  '9',
	"-----":  '0',
	".-.-.-": '.',
	"--..--": ',',
	"..--..": '?',
	".----.": '\u0027', // single quote
	"-..-.":  '/',
	"-.--.":  '(',
	"-.--.-": ')',
	"---...": ':',
	"-...-":  '=',
	".-.-.":  '+',
	"-....-": '-',
	".-..-.": '"',
	".--.-.": '@',
	"..-..":  'é',
}

// the reverse of m2t, to be created in init:
var t2m map[rune]string = map[rune]string{}

func init() {
	for morse, text := range m2t {
		t2m[text] = morse // todo: detect duplicates?
	}
}

func fromMorse(m string) (string, error) {
	if m == "" {
		return "", nil // special case to ease error handling in Decode
	}
	r, ok := m2t[m]
	if !ok {
		return "", ErrUnknownMorse
	}
	return string(r), nil
}

func toMorse(r rune) string {
	if s, ok := t2m[unicode.ToLower(r)]; ok {
		return s
	}
	return "" // silently ignore chars without morse representation
}
