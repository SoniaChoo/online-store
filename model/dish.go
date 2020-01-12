package model

import "time"

type Dish struct {
	DishId    int
	RestId    int
	Price     float64
	DishName  string
	Stock     int
	Sales     int
	CreatTime time.Time
}
