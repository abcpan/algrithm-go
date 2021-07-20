package main

import (
	"fmt"
	"sync"
)



func calcSum(num int, wg *sync.WaitGroup) {
	var sum = 0
	for i := 0; i < 100; i++ {
		sum += i
		fmt.Printf("number %d sum: %d\n",num, sum)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go calcSum(1, &wg)
	go calcSum(2, &wg)
	wg.Wait()
	fmt.Println("计算结束")
}
