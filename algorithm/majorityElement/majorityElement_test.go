package algorithm

import "testing"

type Majority struct {
	Input []int
	Want int
}

func TestMajorityElement(t *testing.T) {
	tests := []Majority{
		{[]int{3,2,3}, 3},
		{[]int{2,2,1,1,1,2,2}, 2},
	}

	for _, tt := range tests {
		ret := majorityElement(tt.Input)
		if ret != tt.Want {
			t.Errorf("error!")
		}
	}
}