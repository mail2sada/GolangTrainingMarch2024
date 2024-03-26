package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Demo: tcp client")
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		fmt.Println("received connection err", err)
	} else {
		conn.Write([]byte("Hello this from our test client..."))
	}
}
