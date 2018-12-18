package main

import (
	"agfun/agfun-service/log"
	"fmt"
	"reflect"
)

func main() {
	url := "/test1/test2/test3"
	if isNil(&url) {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}


func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	log.Warn("Response type should be a pointer")
	return vi.IsValid()
}