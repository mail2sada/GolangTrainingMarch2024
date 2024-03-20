package main

import (
	"first/geo"
	. "first/math"
	adv "first/math/advanced"
	"fmt"

	_ "math"

	"github.com/fatih/color"
	. "github.com/fatih/color"
)

func init() {
	fmt.Println("Init from main")
}

func main() {
	fmt.Println("This is my Package and modules demo...")
	color.Cyan("This is from external package...")
	Cyan("Wonderful....")
	res := Add(1, 2)

	fmt.Println(res)

	fmt.Println(Sub(10, 20))

	fmt.Println(Mul(100, 200))

	fmt.Println(Dic(200, 10))

	fmt.Println(adv.Square(10))

	fmt.Println(geo.AsianContries())

	MethodFromGeo()

	color.Red("This is from ext package")
}
