package maxarea

import "testing"

type MaxArea struct {
	Input []int
	Want int
}

func TestMaxArea(t *testing.T) {
	tests := []MaxArea{
		{[]int {2,3,4,5,18,17,6}, 17},
		{[]int {1,2,1}, 2},
	}
	for _, tt := range tests {
		ret := maxArea(tt.Input)
		if ret != tt.Want {
			t.Errorf("error:")
		}
	}
}