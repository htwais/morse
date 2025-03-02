package convert

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func ExampleEncode() {
	var buf bytes.Buffer
	_ = Encode(&buf, strings.NewReader("SOS"))
	fmt.Println(buf.String())
	// Output: ... --- ...
}

// valid UTF8 must encode to morse without errors:
func TestEncodeValid(t *testing.T) {
	for i, tc := range []struct {
		text     string
		expected string
	}{
		{" ", "/"},
		{"\n", "//"},
		{"e", "."},
		{"The quick brown fox jumps over the lazy dog", "- .... ./--.- ..- .. -.-. -.-/-... .-. --- .-- -./..-. --- -..-/.--- ..- -- .--. .../--- ...- . .-./- .... ./.-.. .- --.. -.--/-.. --- --."},
	} {
		in := strings.NewReader(tc.text)
		var out bytes.Buffer
		if err := Encode(&out, in); err != nil {
			t.Errorf("testcase %d failed with unexpected error: %v", i, err)
		} else if tc.expected != out.String() {
			t.Errorf("testcase %d: expected '%s', got '%s'", i, tc.expected, out.String())
		}
	}
}

// invalid UTF8 must produce an error:
func TestEncodeInvalid(t *testing.T) {
	text := string([]byte{195, 40})
	expectedErr := "invalid UTF8"

	in := strings.NewReader(text)
	var out bytes.Buffer
	err := Encode(&out, in)
	if err == nil || err.Error() != expectedErr {
		t.Fatalf("expected error '%s', got %v", expectedErr, err)
	}
}
