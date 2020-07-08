package main

import "fmt"

func main() {
	greeting := hello("world")
	fmt.Println(greeting)
}

func hello(word string) string {
	return fmt.Sprintf("Hello %v", word)
}
