package main

import "fmt"

func main() {
	fmt.Println("Annonymous with defer")

	i := 0

	defer func(in *int) {
		fmt.Println(*in)
	}(&i)

	i = 100

	fmt.Println("We are exiting main")

}
