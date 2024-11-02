package main
import (
    "testing"
	"unicode/utf8"
)

func TestReserve(t *testing.T) {
	testcases := []struct {
		input, expected string
	}{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
		{"qwerty", "ytrewq"},
	}
	
	for _, tc := range testcases {
		rev, err := ReverseByte(tc.input)
		if rev != tc.expected {
			t.Errorf("Reverse : %q, expected %q", rev, tc.expected)
		}
		if err != nil {
			t.Errorf("General issue : %q", err)
		}
	}
}

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc) // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
            return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}