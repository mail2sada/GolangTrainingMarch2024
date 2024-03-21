package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const FanoutCount = 8

func Generator(list ...string) chan string {
	wg.Add(1)
	gch := make(chan string)

	go func() {
		defer wg.Done()
		defer close(gch)
		for _, v := range list {
			gch <- v
		}
	}()
	return gch
}

func Process(gch <-chan string) chan string {
	wg.Add(1)
	dch := make(chan string)
	fwg := sync.WaitGroup{}
	for i := 0; i < FanoutCount; i++ {
		fwg.Add(1)
		go func() {
			defer fwg.Done()

			for str := range gch {
				rev := ""
				for _, c := range str {
					rev = string(c) + rev
				}

				// Introducing delay
				for i := 0; i < math.MaxUint16; i++ {

				}
				dch <- rev
			}
		}()
	}

	go func() {
		defer wg.Done()
		defer close(dch)
		fwg.Wait()
	}()

	return dch
}

func Display(dch <-chan string) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for str := range dch {
			fmt.Println(str)
		}
	}()
}

func main() {
	defer func(t time.Time) {
		fmt.Println("execution time:", time.Since(t))
	}(time.Now())
	fmt.Println("Demo: Fanout FanIn")

	Display(Process(Generator("Hi", "Hello", "GA", "GM", "Golang", "Erlang", "Elix", "Rust")))
	wg.Wait()
}
