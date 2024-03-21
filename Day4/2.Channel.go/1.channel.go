package main

import "fmt"

var ch = make(chan int)

func Factorial(num int) {
	fact := 1

	for idx := 1; idx <= num; idx++ {
		fact *= idx
	}
	fmt.Println("Calculated factorial for", num, "and writing this to channel")
	ch <- fact
}

func main() {
	fmt.Println("Demo channels...")

	for i := 1; i <= 10; i++ {
		go Factorial(i)
	}
	for i := 0; i < 10; i++ {
		data := <-ch
		fmt.Println(data)
	}
}
