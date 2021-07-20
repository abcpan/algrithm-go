package main

import (
	"fmt"
	"time"
)

func main() {
	// 延时器
	//timer := time.NewTimer(2000 * time.Millisecond)
	//fmt.Println("now: %v", time.Now())
	//t := <- timer.C
	//fmt.Println("t: %v\n", t)

	//timer := time.NewTimer(4000 * time.Millisecond)
	//go func() {
	//	<- timer.C
	//	fmt.Println("四秒钟后执行中....")
	//}()
	var i int
  ticker := SetInterval(func() {
		 fmt.Println("Hello World")
		 i++
 }, 1000)



	for {
		if i == 20 {
			ClearInterval(ticker)
			fmt.Println("定时器执行完毕!")
			break
		}
	}
}

func SetInterval(f func(), tick int64) *time.Ticker {
	duration := time.Duration(tick * 1000 * 1000)
	timer := time.NewTicker(duration)
	go func( ticker *time.Ticker ) {
		for{
			<-ticker.C
			f()
		}
	}(timer)
	return timer
}

func ClearInterval(ticker *time.Ticker) {
	ticker.Stop()
}