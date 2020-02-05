package model

import "time"

type Dish struct {
	DishId      int
	RestId      int
	Price       float64
	DishName    string
	Description string
	Stock       int
	Sales       int
	Favorite    int
	CreatTime   time.Time
	UpdateTime  time.Time
}
