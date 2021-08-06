package floodFill

import (
	"reflect"
	"testing"
)

type FloodFillTest struct {
	Input [][]int
	Want [][]int
}

func TestFloodFill(t *testing.T) {
	tests := []FloodFillTest {
		{[][]int{
			{1,1,1},
			{1,1,0},
			{1,0,1},
		},
		[][]int {
			{2,2,2},
			{2,2,0},
			{2,0,1},
		},
		},
	}
	for _, tt := range tests {
		ret := floodFill(tt.Input, 1, 1, 2)
		if !reflect.DeepEqual(ret, tt.Want) {
			t.Errorf("expected %v got %v", tt.Want, ret)
		}
	}
}