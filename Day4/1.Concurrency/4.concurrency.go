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

func SchedulePrint(msg string) {
	wg.Add(1)
	go Print(msg)
}

func main() {
	fmt.Println("Demo: Wait groups")
	SchedulePrint("This is from my first thread")
	SchedulePrint("This is from my second thred")
	SchedulePrint("This is from my third go routine")
	wg.Wait()
}
