package main

import "fmt"

type MError struct {
	code int
	desc string
}

func (e MError) Error() string {
	return fmt.Sprintf("Received Code:%d %s %s", e.code, "desc:", e.desc)
}

func Division(a, b int) (res int, err error) {
	if b == 0 {
		err = MError{code: 10, desc: "Divide by zero"}
		//err = fmt.Errorf("received error")
		return
	}
	return a / b, nil
}

func main() {
	fmt.Println("Custom error...")

	res, err := Division(100, 0)
	if err != nil {
		cErr := err.(MError)
		fmt.Println(cErr.code, cErr.desc)
		fmt.Println("Received error", err)

	} else {
		fmt.Println(res)
	}

}
