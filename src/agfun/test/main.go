package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := A{
		Name: "xiaohua",
		Age:  15,
	}
	buf, e := json.Marshal(&a)
	if e != nil {
		panic(e)
	}
	var b A
	Get(buf, &b)
	fmt.Println(b)
}

type A struct {
	Name string
	Age  int
}

func Get(buf []byte, value interface{}) {
	json.Unmarshal(buf, value)
}
