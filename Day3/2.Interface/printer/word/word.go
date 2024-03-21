package word

import "fmt"

type Word struct {
	docId   int
	docName string
}

func (w Word) Print() {
	fmt.Println("Printing Word document")
	fmt.Println(w.docId)
	fmt.Println(w.docName)

}
func (w Word) Preview() {
	fmt.Println("Previewing Word document")
}
func (w Word) PageSetup() {

	fmt.Println("PageSetup Word document")

}
