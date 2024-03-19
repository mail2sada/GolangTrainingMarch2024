package main

import "fmt"

func main() {
	fmt.Println("Exploring annonymous structs ")

	as := struct {
		int
		float32
		string
	}{int: 0, float32: 3.142, string: "Hi"}

	fmt.Println(as)

	a := as

	a.int = 100

	fmt.Println(as, a)
}
