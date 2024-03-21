package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func Print(msg string) {
	fmt.Println(msg)
	wg.Done()
}

func main() {
	fmt.Println("Demo: Wait groups")
	wg.Add(3)
	go Print("This is from my first thread")

	go Print("This is from my second thred")

	go Print("This is from my third go routine")
	wg.Wait()
}
