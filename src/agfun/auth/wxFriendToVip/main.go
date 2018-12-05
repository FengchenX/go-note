package main

import (
	"agfun/auth/entity"
	"github.com/jinzhu/gorm"
	"log"
	"math"
	"time"
)

func main() {
	db, e := gorm.Open("mysql",
		"root:feng@tcp(localhost:3306)/auth?charset=utf8&parseTime=true&loc=Local")
	if e != nil {
		log.Fatal(e)
	}
	for {
		time.Sleep(10 * time.Minute)
		var total int
		offset := 0
		num := 10
		db.Model([]entity.WXFriend{}).Count(&total)
		if db.Error != nil {
			log.Fatal(db.Error)
		}
		ceil := int(math.Ceil(float64(total) / float64(num)))
		for i := 0; i <= ceil; i++ {
			var wxFriends []entity.WXFriend
			db := db.Find(&wxFriends).Limit(10).Offset(offset)
			if db.Error != nil {
				log.Fatal(db.Error)
			}
			offset = offset + num

		}

	}
}
