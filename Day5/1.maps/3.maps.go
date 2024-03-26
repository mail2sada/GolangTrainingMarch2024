package main

import "fmt"

func main() {
	fmt.Println("Maps continued")
	mp := make(map[int]int)

	//fmt.Println(mp)

	mp[0] = 1
	mp[2] = 10
	mp[1] = 100
	mp[100] = 1

	for k, v := range mp {
		fmt.Println("key:", k, "value:", v)
	}
}
