package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Exploring strings package...")

	str := "1,2,3,4,5,6,6,7,8"

	values := strings.Split(str, ",")

	for _, v := range values {
		fmt.Println(v)
	}

	str = strings.Join(values, ",")

	fmt.Println(str)
}
