package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("Fs.Walk")

	root := "/Users/sadanandd/Documents"
	fsm := os.DirFS(root)
	fs.WalkDir(fsm, ".", func(path string, d fs.DirEntry, err error) error {

		fmt.Println(path)
		info, err := d.Info()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(info.Size())
		}
		return nil
	})
}
