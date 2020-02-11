package db

import (
	"context"
	"github.com/SoniaChoo/online-store/model"
	"log"
	"time"
	//"github.com/SoniaChoo/online-store/model"
)

const (
	InCartStatus = -1
)

// NewCart is to insert a order info into database when user register
func NewCart(userid int) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}

	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err = db.QueryContext(ctx, "insert orders(user_id, status, total_price) values(?,?,?)", userid, InCartStatus, 0)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}

func UpdatePriceInOrderDetail(d *model.Dish) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}
	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// "and" can be not replaced by ","
	_, err = db.QueryContext(ctx, "update order_detail set price = ? where dish_id = ? and status = ?", d.Price, d.DishId, InCartStatus)
	if err != nil {
		log.Printf("record updating with error %s\n", err.Error())
		return err
	}

	return nil
}
