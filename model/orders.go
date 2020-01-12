package model

import "time"

type Orders struct {
	OrderId   int
	UserId    int
	Status    int
	CreatTime time.Time
}
