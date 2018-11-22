package entity

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	UID  string
	Name string
	Desc string
	URL  string
}

type FreeVideo struct {
	Video
}

type PaidVideo struct {
	Video
}
