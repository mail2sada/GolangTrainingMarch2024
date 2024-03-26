package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Read file")

	f, e := os.Open("/Users/sadanandd/Documents/GitHub/GolangTrainingMarch2024/Day5/3.filehandling/3.file.go")
	if e != nil {
		fmt.Println(e)
		return
	}
	defer f.Close()
	data := make([]byte, 10, 10)
	str := strings.Builder{}
	for {
		n, e := f.Read(data)
		if e != nil {
			fmt.Println(e)
			break
		} else {

			fmt.Println("Read", n, "bytes")
			//str += string(data)
			str.WriteString(string(data))

			//time.Sleep(10 * time.Millisecond)
		}
	}
	fmt.Println(str.String())

}
