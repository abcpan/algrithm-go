package fib

import (
	"fmt"
	"testing"
)

type FibTest struct {
	Input int
	Want int
}

func TestFib(t *testing.T) {

	ret := fib(60)
	fmt.Println(ret)
}
