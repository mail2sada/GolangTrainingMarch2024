package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Exploring strings")

	myString := "Welcome to Go training..."

	for i := 0; i < len(myString); i++ {
		fmt.Println(myString[i])
		fmt.Println(reflect.TypeOf(myString[i]))
	}

	for i, _ := range myString {
		fmt.Println(i)
	}

	fmt.Println("This is my print:", myString[0])
	//myString[0] = 102
}
