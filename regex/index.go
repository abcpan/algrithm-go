package main

import (
	"fmt"
	"regexp"
)

const Text = "My emial is ccmouse@gmail.com 1663071425@qq.com"

func main() {
	matcher, err := regexp.Compile(`([a-zA-Z0-9]+)@[a-zA-Z0-9]+\.[a-z]+`)
	if err != nil {
		panic(err)
	}
	match := matcher.FindAllString(Text, -1)
	subMatch := matcher.FindAllStringSubmatch(Text, -1)
	fmt.Println(match)
	fmt.Println(subMatch)
}
