package main

import (
	"fmt"
	"reflect"
	"time"
)

func PrintDetail(a interface{}, b interface{}) {
	fmt.Println(a)
	fmt.Println(b)
}

func AddUniversal(a interface{}, b interface{}, elem ...interface{}) int {
	sum := a.(int) + b.(int)
	for _, v := range elem {
		sum += v.(int)
	}
	return sum
}

func Add(a interface{}, b interface{}, elements ...interface{}) (res float64) {

	switch a.(type) {
	case int8:
		res += float64(a.(int8))
	case int:
		res += float64(a.(int))
	case float32:
		res += float64(a.(float32))
	case float64:
		res += a.(float64)
	default:
		panic("Unhandled datatype")
	}

	switch b.(type) {
	case int8:
		res += float64(b.(int8))
	case int:
		res += float64(b.(int))
	case float32:
		res += float64(b.(float32))
	case float64:
		res += b.(float64)
	default:
		panic("Unhandled datatype")

	}

	for _, v := range elements {
		switch v.(type) {
		case int8:
			res += float64(v.(int8))
		case int:
			res += float64(v.(int))
		case float32:
			res += float64(v.(float32))
		case float64:
			res += v.(float64)

		default:
			panic("Unhandled datatype")
		}
	}

	return
}

func main() {

	defer func(t time.Time) {
		fmt.Println("Execution time:", time.Since(t))
	}(time.Now())
	fmt.Println("Empty interface")

	PrintDetail("test", 1)
	PrintDetail(1.1, 1)
	PrintDetail([]int{1, 2, 3, 4, 5, 6}, []string{"hello", "hi"})

	//sum := AddUniversal(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, "string")

	sum := Add(1, 1, 10000, 10)

	fmt.Println(sum, reflect.TypeOf(sum))
}
