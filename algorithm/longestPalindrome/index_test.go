package longestPalindrome

import (
	"fmt"
	"testing"
)

type MaxPalindrome struct {
	Input string
	Want string
}

func TestName(t *testing.T) {
	tests := []MaxPalindrome {
		//{"babad", "aba"},
		{"ac", "a"},
	}
	for _, tt := range tests {
		ret := longestPalindrome(tt.Input)
		fmt.Println("ret:" + ret)
		if ret != tt.Want {
			t.Errorf("error")
		}
	}
}