package main

import "fmt"

func main() {
	fmt.Println("Deleting value from map")

	mp := make(map[int]int)

	mp[0] = 1
	mp[1] = 2
	mp[2] = 3
	mp[3] = 4

	fmt.Println(mp)

	delete(mp, 2)

	fmt.Println(mp)

	mp[1] = 100
	fmt.Println(mp)
}
