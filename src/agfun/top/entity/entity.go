package entity

import (
	"time"
)

type Banner struct {
	ID string `json:"id"`
	VideoID string `json:"video_id"`
	CreateTime time.Time `json:"create_time"`
}
