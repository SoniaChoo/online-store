package model

import "time"

type Orders struct {
	OrderId    int
	UserId     int
	Status     int
	TotalPrice float64
	CreatTime  time.Time
	UpdateTime time.Time
}
