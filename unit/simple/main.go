package main

import "fmt"

func main() {
	greeting := hello("World")
	fmt.Println(greeting)
}

func hello(s string) string {
	return fmt.Sprintf("Hello %v", s)
}
