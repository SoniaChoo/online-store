package model

import "time"

type Rest struct {
	RestId    int
	UserId    int
	Phone     string
	RestName  string
	Address   string
	CreatTime time.Time
}
