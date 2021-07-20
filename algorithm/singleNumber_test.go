package algorithm

import "testing"

type SingleNum struct {
	Input []int
	Want int
}

func TestSingleNumber(t *testing.T){
	tests := []SingleNum{
		{Input:[]int {2,2,1},Want: 1},
		{Input:[]int {4,1,2,1,2},Want: 4},
	}
	for _, tt := range tests {
		ret := singleNumber(tt.Input)
		if ret != tt.Want {
			t.Errorf("error")
		}
	}
}