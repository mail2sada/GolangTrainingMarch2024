package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 3)

var isChClosed = false

var wg = sync.WaitGroup{}

var wg1 = sync.WaitGroup{}

func Factorial(num int) {
	defer wg.Done()
	fact := 1

	for idx := 1; idx <= num; idx++ {
		fact *= idx
	}
	//fmt.Println("Calculated factorial for", num, "and writing this to channel")
	ch <- fact
}

func ChannleMonitor() {
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for !isChClosed {
			fmt.Println(cap(ch), len(ch))
			time.Sleep(100 * time.Nanosecond)
		}
	}()
}

func main() {
	defer func(t time.Time) {
		fmt.Println("Took", time.Since(t))
	}(time.Now())
	fmt.Println("Demo channels...")

	ChannleMonitor()
	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go Factorial(i)
		}
		wg.Wait()
		close(ch)
		isChClosed = true
	}()

	for i := 0; i < 11; i++ {
		_, ok := <-ch
		if !ok {
			break
		}
		//fmt.Println(data)
	}
	wg1.Wait()

}
