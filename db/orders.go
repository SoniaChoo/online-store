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

func GetOrderIdInTableOrder(o *model.Orders) ([]*model.Orders, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from orders where user_id = ?", o.UserId)
	if err != nil {
		log.Printf("record search orders by userid = %d with error %s\n", o.UserId, err.Error())
		return nil, err
	}
	defer row.Close()

	orderids := []*model.Orders{}
	for row.Next() {
		temp := &model.Orders{}
		if err = row.Scan(&temp.OrderId, &temp.UserId, &temp.Status, &temp.TotalPrice, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("record search orders by userid =  %d in loop with error %s\n", o.OrderId, err.Error())
			return nil, err
		}
		orderids = append(orderids, temp)
	}
	return orderids, nil
}

func ShowCartOrder(order_id int) ([]*model.Order_detail, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from order_detail where order_id = ? and status = ?", order_id, InCartStatus)
	if err != nil {
		log.Printf("record search order_detail by orderid = %d with error %s\n", order_id, err.Error())
		return nil, err
	}
	defer row.Close()

	carts := []*model.Order_detail{}
	for row.Next() {
		temp := &model.Order_detail{}
		if err = row.Scan(&temp.DetailId, &temp.RestId, &temp.OrderId, &temp.DishId, &temp.Price, &temp.Number, &temp.Status); err != nil {
			log.Printf("record search order_detail by orderid =  %d in loop with error %s\n", order_id, err.Error())
			return nil, err
		}
		carts = append(carts, temp)
	}
	return carts, nil
}

//update total_price in table order when showcart
func UpdateTotalPriceInOrder(total_price float64, order_id int) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "update orders set total_price = ? where order_id = ? and status = ?", total_price, order_id, InCartStatus)
	if err != nil {
		log.Printf("record update total_price in table order by orderid = %d with error %s\n", order_id, err.Error())
		return err
	}
	return nil
}
