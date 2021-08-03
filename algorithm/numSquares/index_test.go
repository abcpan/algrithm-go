package numSquares

import (
	"testing"
)

type NumSquareTest struct {
	Input int
	Want int
}

func TestNumsSquare(t *testing.T) {
	tests := []NumSquareTest {
		{12, 3},
		{13, 2},
	}
	for _, tt := range tests {
		ret := numSquares(tt.Input)
		if ret != tt.Want {
			t.Errorf("expected %v, got %v", tt.Want, ret)
		}
	}
}