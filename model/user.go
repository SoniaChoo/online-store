package model

import "time"

type User struct {
	UserId    int
	Phone     string
	Nickname  string
	Password  string
	CreatTime time.Time
}
