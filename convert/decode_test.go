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
