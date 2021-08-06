package evalRPN

import "testing"

type EvalTest struct {
	Input []string
	Want int
}

func TestEvalRPN(t *testing.T) {
	tests := []EvalTest {
		{[]string{"4","13","5","/","+"}, 6},
		{[]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}, 22},
		{[]string{"2","1","+","3","*"}, 9},
	}
	for _, tt := range tests {
		ret := evalRPN(tt.Input)
		if ret != tt.Want {
			t.Errorf("expected %v, got %v", tt.Want, ret)
		}
	}
}