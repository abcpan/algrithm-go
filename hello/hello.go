package main

import (
	"abcpan.net.cn/greetings"
	"abcpan.net.cn/hello/util"
	"fmt"
	hello_go "github.com/abcpan/hello-go"
	"log"
)
//import "abcpan.net.cn/greetings"

func main() {
	log.SetPrefix("greetings: ")
	// 禁用打印
	log.SetFlags(0)
	msg, err := greetings.Hello("Sli")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

	memo := map[string] int {
		"a": 1,
		"b": 2,
		"c": 3,
	}

	changeMemo(&memo)
	fmt.Println(memo)
	name := util.GetName("pan")
	fmt.Println(name)
	fmt.Println(util.BinarySearch("3333"))
	message := hello_go.Greeting("pan111")
	fmt.Println(message)
}

func changeMemo(memo *map[string] int) {
	(*memo)["a"] = 22
}