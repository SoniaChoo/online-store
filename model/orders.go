package model

import "time"

type Orders struct {
	OrderId   int
	UserId    int
	Status    int
	Price     float64
	CreatTime time.Time
}
