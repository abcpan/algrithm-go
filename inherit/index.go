package main

import "fmt"

type Animal struct {
	color string
	name string
}

type Dog struct {
	*Animal
}

func (dog *Dog) move() {
	fmt.Println(dog.name + "is move")
}


func main() {
	dog := &Dog{
		&Animal{
			name: "旺财",
			color: "黄色",
		},
	}

	dog.move()
}