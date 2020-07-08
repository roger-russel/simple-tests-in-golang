package main

import "fmt"

func main() {
	m := getSlice()
	fmt.Println(m)
}

func getSlice() []string {
	return []string{
		"one",
		"three",
		"two",
	}
}
