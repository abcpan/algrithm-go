package dayliyTemperature

import (
	"fmt"
	"testing"
)

type Temperature struct {
	Input []int
	want [] int
}

func TestDailyTemperature(t *testing.T) {
	tests := []Temperature {
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}
	for _, tt := range tests {
		ret := dailyTemperatures(tt.Input)
		fmt.Println(ret)
	}
}