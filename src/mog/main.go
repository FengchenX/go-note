package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)



func main() {
	url := "localhost:27017"
	session, e := mgo.Dial(url)
	if e != nil {
		panic(e)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("inventory")
	var results []Inventory
	selector := bson.M{}
	e = c.Find(selector).All(&results)
	if e != nil {
		log.Fatal(e)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
	count, e := c.Find(selector).Count()
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(count)

	fmt.Println("********************")
	selector = bson.M{
		"size.h": 19,
	}
	e = c.Find(selector).All(&results)
	if e != nil {
		log.Fatal(e)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

type Inventory struct {
	Item string
	Qty  int
	Tags []string
	Size Size
}
type Size struct {
	H   int
	W   int
	Uom string
}


