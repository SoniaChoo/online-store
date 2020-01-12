package model

import "time"

type Picture struct {
	PicId int
	DishId int
	UserId int
	RestId int
	PicPath string
	FirstPic int
	CreatTime time.Time
}
