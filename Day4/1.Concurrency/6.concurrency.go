package main

/*

generate numbers from 1 .. 10

calculate factorial of this...
*/
import (
	"fmt"
	"sync"
)

func Factorial(num int) {
	defer wg.Done()
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println()
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Factorial(i)
	}
	wg.Wait()
}
