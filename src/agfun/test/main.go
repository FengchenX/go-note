package main

import (
	"agfun/util"
	"encoding/json"
	"fmt"
)

func main() {
	a := A{
		ID:   "aid",
		Name: "myname",
	}
	b := B{
		ID:  "bid",
		Age: 10,
	}
	var c C
	util.Copy(&c.A, &a)
	fmt.Println(c)
	util.Copy(&c.B, &b)
	fmt.Println(c)
	buf, _ := json.Marshal(&c)
	fmt.Println(string(buf))
}

type A struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type B struct {
	ID  string `json:"id"`
	Age int    `json:"age"`
}
type C struct {
	A
	B
}
