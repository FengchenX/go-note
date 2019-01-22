package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "test1/test2/test3"

	if strings.Contains(url, "") {
		fmt.Println("c")
	} else {
		fmt.Println("n")
	}
}


