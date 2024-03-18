package main

import "fmt"

/*
=
+=
-=
*=
%=
&=
|=
*/

func main() {
	fmt.Println("Demo: Assignment ops")

	var a, b = 10, 20

	a = a + b

	fmt.Println(a) // 30

	a += b

	fmt.Println(a) // 50

}
