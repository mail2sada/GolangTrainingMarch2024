package main

import "fmt"

func main() {
	fmt.Println("Demo: Loop and go to ")

outer_loop:
	for i := 1; i < 100; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(j)
			if j%8 == 0 {
				goto outer_loop
			}

		}
	}

}
