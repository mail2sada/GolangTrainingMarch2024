package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Exploring concurrency")

	go fmt.Println("This is my goroutine print")

	fmt.Println("Exiting main")

	go fmt.Println("This is my second go routine...")
	time.Sleep(1 * time.Nanosecond)
}
