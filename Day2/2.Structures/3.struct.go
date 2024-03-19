package main

import "fmt"

type PromptedFields struct {
	int
	uint8
	string
	float32
	bool
	a int
}

func main() {
	fmt.Println("Prompted fields..")

	var pf = PromptedFields{int: 0, uint8: 1, string: "hello", float32: 3.14159, bool: true}

	fmt.Println(pf)

	pf.int = 10000

	pf.string = "Go traing..."

	fmt.Println(pf)
}
