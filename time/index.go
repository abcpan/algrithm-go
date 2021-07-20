package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	ret := time.Unix(now.Unix(), 0)
	fmt.Println(ret)
 fmt.Println(now.Unix())
}
