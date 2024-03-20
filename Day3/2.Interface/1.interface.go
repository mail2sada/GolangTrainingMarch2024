// /*
// type <Name Of Interface> interface{
// 	Method1()
// 	Method2()
// }
// */

// package main

// import "fmt"

// type Emp struct {
// 	Id   int
// 	Name string
// }

// func (e Emp) Print() {
// 	fmt.Println("eID:", e.Id)
// 	fmt.Println("eName:", e.Name)
// }

// func (e Emp) String() string {
// 	fmt.Println("String in emp")
// 	return fmt.Sprint("eID:", e.Name, "eName:", e.Name)
// }

// type Std struct {
// 	Id   int
// 	Name string
// }

// func (s Std) Print() {
// 	fmt.Println("sID:", s.Id)
// 	fmt.Println("sName:", s.Name)
// }

// func (s Std) String() string {
// 	fmt.Println("String() int std")
// 	return fmt.Sprint("sID:", s.Name, "sName:", s.Name)
// }

// type Printer interface {
// 	Print()
// }

// func Print(p Printer) {
// 	p.Print()
// }

// func main() {
// 	fmt.Println("Demo: Interface")
// 	var e, s Printer

// 	e = Emp{Id: 1, Name: "Anand"}

// 	s = Std{Id: 1, Name: "Anilkumar"}

// 	fmt.Println(e)

// 	fmt.Println(s)
// }
