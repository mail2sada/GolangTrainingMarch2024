package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	RollNo int    `json:"roll_no"`
	Name   string `json:"name"`
	Class  string `json:"class"`
}

func main() {
	fmt.Println("Json encoding...")

	std := Student{RollNo: 1, Name: "Anand", Class: "11B"}

	fmt.Println(std)

	data, err := json.Marshal(std)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	data, err = json.MarshalIndent(std, "\n", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
