package bmp

import "fmt"

type Bmp struct {
}

func (b Bmp) Print() {
	fmt.Println("Printing bitmap")
}

func (b Bmp) PageSetup() {
	fmt.Println("Pagesetup bitmap")
}

func (b Bmp) Preview() {
	fmt.Println("Previwing bitmap")
}
