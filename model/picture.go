package model

import "time"

type Picture struct {
	PicId     int
	Types     int
	PicPath   string
	FirstPic  int
	CreatTime time.Time
}
