package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Pointers...")

	var ar [3]string = [3]string{"Hi", "Hello", "How are you"}

	ptrArray := &ar //ptrArray is pointer string array

	ptrEle1 := &ar[0]

	fmt.Println(ptrArray)
	fmt.Println(*ptrArray)
	fmt.Println(ptrEle1)

	*ptrEle1 = "This is wonderful"

	fmt.Println(*ptrEle1)
	fmt.Println(reflect.TypeOf(ptrEle1))
	fmt.Println(reflect.TypeOf(ptrArray))

	for i, v := range ptrArray {
		fmt.Println(i, v)
		fmt.Println(reflect.TypeOf(v))
	}

}
