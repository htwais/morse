package convert

import (
	"bytes"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	for i, tc := range []struct {
		morse    string
		expected string
	}{
		{".", "e"},
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
