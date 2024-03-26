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
	fmt.Println("Json Encoding...")

	jsonData := `{

        "roll_no": 1,

        "name": "Anand",

        "class": "11B"

}`
	fmt.Println(jsonData)
	data := []byte(jsonData)
	fmt.Println(data)
	std := Student{}
	err := json.Unmarshal(data, &std)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(std)
		fmt.Printf("%+v", std)
	}
}
