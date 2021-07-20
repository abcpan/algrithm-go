package generateParenthesis

import (
	"reflect"
	"testing"
)

type ParenthesisTest struct {
	Input int
	Want []string
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []ParenthesisTest {
		{2, []string{"(())","()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, tt := range tests {
		ret := generateParenthesis(tt.Input)
		if !reflect.DeepEqual(ret, tt.Want){

			t.Errorf("error!, %v", ret)
		}
	}
}