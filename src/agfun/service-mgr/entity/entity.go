package entity

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Name string
	Desc string
	URL string
}