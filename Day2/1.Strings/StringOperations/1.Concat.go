package main

import "fmt"

func main() {
	fmt.Println("Working string operations...")

	a, b := "first string", `second string`

	c := a + b

	i := 100

	fmt.Println(c)

	c = fmt.Sprint(a, b)
	fmt.Println(c)

	c = fmt.Sprintln("FirstString:", a, "SecondString:", b, i)

	fmt.Println(c)

}
