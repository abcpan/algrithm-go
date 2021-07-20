package phoneNumber

import (
	"fmt"
	"reflect"
	"testing"
)

type PhoneNumber struct {
	Input string
	Want [] string
}

func TestPhoneNumber(t *testing.T) {
	tests := []PhoneNumber{
		{"23", []string{"ad","ae","af","bd","be","bf","cd","ce","cf"}},
		{"2", []string{"a","b","c"}},
	}

	for _, tt := range tests {
		ret := letterCombinations(tt.Input)
		fmt.Println(ret)
		if !reflect.DeepEqual(ret, tt.Want) {
			t.Errorf("error!")
		}
	}
}