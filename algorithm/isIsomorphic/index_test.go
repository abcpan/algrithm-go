package isIsomorphic

import "testing"

type OmorphicTest struct {
	S1 string
	S2 string
	Want bool
}

func TestIsomorphic(t *testing.T) {
	tests := []OmorphicTest {
		{"aa", "ab", false},
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}

	for _, tt := range tests {
		ret := isIsomorphic(tt.S1, tt.S2)
		if ret != tt.Want {
			t.Errorf("error! %v, %v", tt.S1, tt.S2)
		}
	}
}
