package db

import (
	"context"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"
)

// InsertOrderDetail is to insert a order detail info into database
func InsertOrderDetail(od model.Order_detail) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}
	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into user(detail_id, dish_id, rest_id, order_id, price, number) values(?,?,?,?,?,?)",
		od.DetailId, od.DishId, od.RestId, od.OrderId, od.Price, od.Number)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}