package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Handling timers")
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		defer wg.Done()
		fmt.Println("Timer is expired...")
	})

	wg.Wait()

}
