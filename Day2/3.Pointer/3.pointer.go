package main

import "fmt"

type comp struct {
	ModelNo string
	Name    string
}

func main() {
	fmt.Println("Pointers...")

	v := comp{ModelNo: "12232BHH", Name: "Hp Pavillion"}

	fmt.Println(v)

	fmt.Println(v.ModelNo, v.Name)

	ptrComp := &v

	fmt.Println(ptrComp)

	fmt.Println(ptrComp.ModelNo, ptrComp.Name)

	var ptrC *comp = &v

	fmt.Println(ptrC)
}
