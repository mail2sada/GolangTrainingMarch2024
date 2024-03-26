package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading a file...")
	data, err := os.ReadFile("/Users/sadanandd/Documents/GitHub/GolangTrainingMarch2024/Day5/3.filehandling/1.fileread.go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
