package convert

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func ExampleDecode() {
	var buf bytes.Buffer
	_ = Decode(&buf, strings.NewReader("... --- ..."))
	fmt.Println(buf.String())
	// Output: sos
}

func TestDecode(t *testing.T) {
	for i, tc := range []struct {
		morse    string
		expected string
	}{
		{"", ""},
		{"/", " "},
		{"//", "\n"},
		{"///", "\n "},
		{".", "e"},
		{".--. . .-. -.-. .... ..-../..-..", "perché é"}, // é is "part of the ITU-R Morse code standard"
		{"--..--/.--", ", w"},
	} {
		in := strings.NewReader(tc.morse)
		var out bytes.Buffer
		if err := Decode(&out, in); err != nil {
			t.Errorf("testcase %d failed with unexpected error: %v", i, err)
		} else if tc.expected != out.String() {
			t.Errorf("testcase %d: expected '%s', got '%s'", i, tc.expected, out.String())
		}
	}
}

func TestDecodeFail(t *testing.T) {
	for i, tc := range []struct {
		morse    string
		expected string
	}{
		{"x", "invalid morse"},
		{"-.-.--", "unknown morse"}, // nonstandard punctuation '!'
	} {
		var out bytes.Buffer
		if err := Decode(&out, strings.NewReader(tc.morse)); err == nil {
			t.Errorf("testcase %d failed: expected error '%s', got none", i, tc.expected)
		} else if tc.expected != err.Error() {
			t.Errorf("testcase %d failed: expected error '%s', got %v", i, tc.expected, err)
		}
	}
}
