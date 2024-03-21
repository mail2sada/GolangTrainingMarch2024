package main

import (
	"fmt"
	"printer/bmp"
	"printer/print"
	"printer/word"
)

func main() {
	fmt.Println("Demo interfaces")

	var p print.Printer

	p = bmp.Bmp{}

	print.Print(p)
	print.PageSetup(p)
	print.Preview(p)

	var p1 print.Printer = word.Word{}
	print.Print(p1)
	print.PageSetup(p1)
	print.Preview(p1)
}
