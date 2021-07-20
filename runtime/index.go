package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(msg string) {
		for i := 0; i < 20; i++ {
			fmt.Println(msg)
		}
	}("World")

	for i := 0; i < 20; i++ {
		runtime.Gosched()
		fmt.Println("Hello")
	}
}
