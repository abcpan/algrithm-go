package isHappy

import "testing"

type HappyTest struct {
	Input int
	Want bool
}

func TestIsHappy(t *testing.T) {
	tests := []HappyTest {
		{7, true},
		{2, false},
	}

	for _, tt := range tests {
		ret := isHappy(tt.Input)
		if ret != tt.Want {
			t.Errorf("expected %v, got %v",tt.Want, ret)
		}
	}
}