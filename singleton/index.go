package main

import (
	"fmt"
	"sync"
)

var createOnce sync.Once

type Person struct {
	Name string
	Age int
}

var person *Person

func createInstance() {
	person = &Person{
		Name: "Hello",
		Age: 22,
	}
}

func getInstance() *Person {
	createOnce.Do(createInstance)
	return person
}


func main() {
	var limit = 2000
	var wg sync.WaitGroup
	wg.Add(limit)
	for i := 0; i < limit; i++ {
		go func() {
			p := getInstance()
			fmt.Println(&p.Name)
			wg.Done()
		}()
	}
	wg.Wait()
}
