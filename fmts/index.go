package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "小明, 小红, 小华都有光明的未来 %s", "!");
	err := fmt.Errorf("this is a error %s", "/pan/guicheng")
	fmt.Println(err)
}
