package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sort package...")
	v := []int{1, 6, 4, 3, 5, 2}

	fmt.Println(v)
	sort.Ints(v)

	fmt.Println(v)

	a := sort.SearchInts(v, 6)
	fmt.Println(a)

}
