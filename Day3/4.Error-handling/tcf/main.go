package main

import "fmt"

type Tcf struct {
	Try      func()
	Catch    func(interface{})
	Finalize func()
}

func (t Tcf) Do() {
	if t.Finalize != nil {
		defer t.Finalize()
	}
	if t.Catch != nil {
		defer func() {
			a := recover()
			if a != nil {
				t.Catch(a)
			}
		}()
	}
	t.Try()
}

func Divide(a, b int) int {
	return a / b
}

func main() {
	fmt.Println("Simulating try catch...")

	Tcf{
		Try: func() {
			res := Divide(10, 1)
			fmt.Println(res)
			res = Divide(100, 0)
			fmt.Println(res)
		},
		Catch: func(e interface{}) {
			fmt.Println(e)
		},
		Finalize: func() {
			fmt.Println("Happy exit...")
		},
	}.Do()
}
