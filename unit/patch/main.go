package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/tmp/foo/boo"
	err := os.MkdirAll(path, 0755)

	if err != nil {
		fmt.Println(err)
	}
}
