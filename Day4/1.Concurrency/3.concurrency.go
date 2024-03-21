package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func Print(msg string) {
	wg.Add(1)
	fmt.Println(msg)
	wg.Done()
}

func main() {
	fmt.Println("Demo: Wait groups")

	go Print("This is from my first thread")
	wg.Add(1)
	go Print("This is from my second thred")
	wg.Add(1)
	go Print("This is from my third go routine")
	wg.Wait()
}
