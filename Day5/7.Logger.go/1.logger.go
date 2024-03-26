package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Demo: Logger")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetPrefix("[Go training]")

	log.Println("Entering main...")

	//log.Fatalln("Attempting fatalln")

	log.Panic("hello....")

	fmt.Println("Hello..")
}
