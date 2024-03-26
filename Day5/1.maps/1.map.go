package main

import "fmt"

func main() {
	fmt.Println("Working with maps")
	var Map map[string]string = map[string]string{"WRG_URL": "https://wrg.mavenir.com",
		"MSTORE_URL": "https://mstore.mavenir.com",
		"MEDIA_URL":  "media.mavenir.com"}
	for key, value := range Map {
		fmt.Println("Key:", key, "\tValue:", value)
	}
}
