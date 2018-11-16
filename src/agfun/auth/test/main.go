package main

import "fmt"

func main() {
	var a A
	if a.B == nil || a.B.Name != a.Name {
		fmt.Println("return")
	}
}

type A struct {
	B *B
	Name string
}

type B struct {
	Name string
	Age int
}
