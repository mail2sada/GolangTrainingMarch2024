package main

import "fmt"

type Object struct {
	Id   int
	Desc string
}

func (o *Object) Read() {
	fmt.Println(o)
}

func (o *Object) Write(id int, desc string) {
	o.Id = id
	o.Desc = desc
}

type Reader interface {
	Read()
}

type Writer interface {
	Write(id int, desc string)
}

type ObjectIO interface {
	Reader
	Writer
}

func main() {
	fmt.Println("Embedded interface")

	var r Reader

	var w Writer

	var oIO ObjectIO

	oIO = &Object{}

	oIO.Write(1, "Test Object")
	oIO.Read()

	r = oIO

	w = oIO
	w.Write(1, "Dummy object")
	r.Read()

}
