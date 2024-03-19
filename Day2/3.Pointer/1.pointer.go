package main

import "fmt"

func main() {
	fmt.Println("Pointers....")

	i := 100

	var ptrInt *int = &i

	fmt.Println("Value of:", i) //100

	fmt.Println("Address of:", &i) // 0x00001234

	fmt.Println("Contents of pointer", ptrInt) //0x00001234

	fmt.Println("Contents of int using pointer", *ptrInt) //100
}
