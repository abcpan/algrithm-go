package subsets

import (
	"fmt"
	"testing"
)

type BackTracking struct {
	Input []int
	Want [][]int
}

func TestSubsets(t *testing.T) {
	tests := []BackTracking{
		{[]int {1,2,3}, [][]int {{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		ret := subsets(tt.Input)
		fmt.Println(ret)
	}
}