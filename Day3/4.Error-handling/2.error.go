package main

import (
	"fmt"
)

func Div(a, b int) (int, error) {
	if b == 0 {
		// this a panic case
		// create an error and return
		//err := errors.New("We have received an error...")
		err := fmt.Errorf("Received erro as denominater is %d", b)
		return 0, err
	}
	return a / b, nil
}

func main() {

	res, err := Div(10, 1)

	if err != nil {
		fmt.Println("Received error...", err)
	} else {
		fmt.Println("Result:", res)
	}
	res, err = Div(10, 0)
	if err != nil {
		fmt.Println("Received error...", err)
	} else {
		fmt.Println("Result:", res)
	}
}
