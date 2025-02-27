package convert

import (
	"bytes"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	for i, tc := range []struct {
		text     string
		expected string
	}{
		{"e", "."},
	} {
		in := strings.NewReader(tc.text)
		var out bytes.Buffer
		if err := Decode(&out, in); err != nil {
			t.Errorf("testcase %d failed with unexpected error: %v", i, err)
		} else if tc.expected != out.String() {
			t.Errorf("testcase %d: expected '%s', got '%s'", i, tc.expected, out.String())
		}
	}
}
