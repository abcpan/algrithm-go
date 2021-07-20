package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	str := "100233a"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Println(num)
	}else if err != nil {
		log.Fatal(err)
	}
}
