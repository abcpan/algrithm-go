package calc

import "testing"

type MaxProfit struct {
	Input [] int
	Want int
}

func TestMaxProfit(t *testing.T) {
	tests  := []MaxProfit{
		{[]int{7,1,5,3,6,4}, 5},
		{[]int{7,6,4,3,1}, 0},
	}
	for _, tt := range tests {
		ret := maxProfit(tt.Input)
		if ret != tt.Want {
			t.Errorf("error!")
		}
	}
}