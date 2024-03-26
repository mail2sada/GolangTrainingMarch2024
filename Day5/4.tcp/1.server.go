package main

import (
	"fmt"
	"log"
	"net"
)

func HandleConnection(c net.Conn) {
	defer c.Close()
	data := make([]byte, 1024, 1024)
	n, err := c.Read(data)
	if err != nil {
		fmt.Println("Received error while reading", err)
	} else {
		fmt.Println("Read ", n, "bytes", "Contents:", string(data))
	}
}
func main() {
	fmt.Println("Demo: tcp server")
	l, e := net.Listen("tcp", ":1234")

	if e != nil {
		log.Fatalln("Failed to listen", e)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			fmt.Println("Failed to accept connection", e)
		} else {
			go HandleConnection(c)
		}
	}
}
