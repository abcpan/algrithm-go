package main

import (
	"encoding/json"
	"fmt"
)

type Pet struct {
	Name string
}

type Person struct {
	Pet Pet
}



func main() {
  person := Person{
  	Pet: Pet{"旺财"},
	}

	person.Pet.Name = "看门大将"
	bytes, _ := json.Marshal(person)
	fmt.Println(string(bytes))
}
