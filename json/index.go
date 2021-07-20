package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Roles []string `json:"roles"`
}

func main() {
	// 序列化
	//person := &Person{"潘", 11}
	//bytes, err := json.Marshal(person)
	//if err != nil {
	//	fmt.Println("error !")
	//}
	//var jsonstr = string(bytes)
	//fmt.Println(jsonstr)
	//
	//// 反序列化
	//var newPerson Person
	//fmt.Println(newPerson)
	//if err := json.Unmarshal([]byte(jsonstr), &newPerson); err == nil {
	//	fmt.Println(newPerson.Name)
	//}

	// 测试空类型
	var person = Person{}
	fmt.Println(person)
	bytes, err := json.Marshal(person)
	if err == nil {
		fmt.Println(string(bytes))
	}
}
