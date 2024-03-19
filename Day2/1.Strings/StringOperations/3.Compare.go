package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Exploring strings package")

	str1, str2 := "string2", "string2"

	if str1 == str2 {
		fmt.Println("Strings are equal")
	}

	if strings.Compare(str1, str2) == 0 {
		fmt.Println("strings are equal")
	}
	fmt.Println("String Compare output:", strings.Compare(str1, str2))
	if reflect.DeepEqual(str1, str2) {
		fmt.Println("Strings are equal")
	}
}
