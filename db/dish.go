package db

import (
	"context"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"
)

// InsertDish is to insert a dish info into database
func InsertDish(d model.Dish) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}
	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into user(dish_id, rest_id, price, dish_name, stock, slaes, creat_time) values(?,?,?,?,?,?,?)",
		d.DishId, d.RestId, d.Price, d.DishName,d.Stock,d.Sales,d.CreatTime)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}