package main

import "fmt"

func main() {
	fmt.Println("Explorining maps")

	var mp map[string]int

	mp["Hello"] = 1  // core
	k := mp["Hello"] // this wil not
	fmt.Println(mp, k)
}
