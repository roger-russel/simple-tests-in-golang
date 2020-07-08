package main

import "fmt"

//Obj is just a sample obj
type Obj struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	obj := getObj(1, "name", "type")
	fmt.Println(obj)
}

func getObj(id int, name, otype string) Obj {
	return Obj{
		ID:   id,
		Name: name,
		Type: otype,
	}
}
