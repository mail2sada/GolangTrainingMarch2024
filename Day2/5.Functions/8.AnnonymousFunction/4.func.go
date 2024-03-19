package main

import "fmt"

func HigherOrderFunction(str string, fn func(bt rune)) {
	for _, v := range str {
		fn(v)
	}
}

func RuneReader(r rune) {
	fmt.Println(r)
}

func main() {
	fmt.Println("Higher order function")

	var str = "This example is from go traing @mavenir"
	HigherOrderFunction(str, RuneReader)

	anfunc := func(r rune) {
		fmt.Println("Received ", r)
	}

	HigherOrderFunction(str, anfunc)

	HigherOrderFunction(str, func(r rune) {
		fmt.Printf("We just got %c\n", r)
	})
}
