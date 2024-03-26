package main

import "fmt"

func main() {
	fmt.Println("Resetting maps")

	mp := map[int]int{0: 0, 1: 1, 2: 2, 3: 3}

	fmt.Println(mp)

	// for k, _ := range mp {
	// 	delete(mp, k)
	// }

	mp = make(map[int]int)

	fmt.Println(mp)
}
