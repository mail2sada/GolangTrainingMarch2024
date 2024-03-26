package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File write")
	fname := "/Users/sadanandd/Documents/GitHub/GolangTrainingMarch2024/Day5/test.txt"
	err := os.WriteFile(fname, []byte("This is our test content"), 0777)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	data, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
