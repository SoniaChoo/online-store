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

func GetOrderIdInTableOrder(userId int) ([]*model.Orders, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from orders where user_id = ? and status = ?", userId, InCartStatus)
	if err != nil {
		log.Printf("record search orders by userid = %d with error %s\n", userId, err.Error())
		return nil, err
	}
	defer row.Close()

	orderids := []*model.Orders{}
	for row.Next() {
		temp := &model.Orders{}
		if err = row.Scan(&temp.OrderId, &temp.UserId, &temp.Status, &temp.TotalPrice, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("record search orders by userid =  %d in loop with error %s\n", userId, err.Error())
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
	row, err := db.QueryContext(ctx, "select * from order_detail where order_id = ?", order_id)
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
	_, err = db.QueryContext(ctx, "update orders set total_price = ? where order_id = ?", total_price, order_id)
	if err != nil {
		log.Printf("record update total_price in table order by orderid = %d with error %s\n", order_id, err.Error())
		return err
	}
	return nil
}

//add dish to cart
func InsertDishToCart(detail *model.Order_detail) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into order_detail(rest_id, order_id, dish_id, price, number, status) Values(?,?,?,?,?,?)", detail.RestId, detail.OrderId, detail.DishId, detail.Price, detail.Number, InCartStatus)
	if err != nil {
		log.Printf("record insert into table order_detail with error, error is %s\n", err.Error())
		return err
	}
	return nil
}

func UpdateDishToCart(detail *model.Order_detail) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	if _, err = db.QueryContext(ctx, "update order_detail set number = ? where detail_id = ?", detail.Number, detail.DetailId); err != nil {
		log.Printf("record update dish number in table order_detail with error, error is %s\n", err.Error())
		return err
	}
	return nil
}

//update dish favorite when add dish to cart
func UpdateDishFavorite(newfavorite, dish_id int) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	//in the SQL statement, "," can not be replaced by "and"
	_, err = db.QueryContext(ctx, "update dish set favorite = favorite + ? where dish_id = ? ", newfavorite, dish_id)
	if err != nil {
		log.Printf("record update dish with error %s\n", err.Error())
		return err
	}

	return nil
}

func AddToCartOrder(detail *model.Order_detail) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	//search if cart has this dish
	row, err := db.QueryContext(ctx, "select * from order_detail where dish_id = ? and order_id = ?", detail.DishId, detail.OrderId)
	if err != nil {
		log.Printf("record select by dish_id and order_id in table order_detail with error, error is %s\n", err.Error())
		return err
	}

	carts := []*model.Order_detail{}
	for row.Next() {
		temp := &model.Order_detail{}
		if err = row.Scan(&temp.DetailId, &temp.RestId, &temp.OrderId, &temp.DishId, &temp.Price, &temp.Number, &temp.Status); err != nil {
			log.Printf("record select by dish_id = %d and order_id = %d in table order_detail  in loop with error %s\n", detail.DishId, detail.OrderId, err.Error())
			return err
		}
		carts = append(carts, temp)
	}

	//if len(carts) = 0, insert to order_detail, else, update dish number according dish_id
	if len(carts) > 1 {
		log.Printf("lengths of carts should be less than 1, but now it's length is %d\n", len(carts))
		return nil
	}

	//before insert or update, update dish favorite when add dish to cart
	if err = UpdateDishFavorite(detail.Number, detail.DishId); err != nil {
		log.Printf("record update dish favorite in table dish when add dish to cart with error, error is %s\n", err.Error())
		return err
	}

	if len(carts) == 0 {
		if err = InsertDishToCart(detail); err != nil {
			log.Printf("record insert into table order_detail with error, error is %s\n", err.Error())
			return err
		}
		return nil
	}

	//brfore upadte dish number, we need to get the original number of this dish in cart
	detail.Number += carts[0].Number
	detail.DetailId = carts[0].DetailId
	if err = UpdateDishToCart(detail); err != nil {
		log.Printf("record update dish number in table order_detail with error, error is %s\n", err.Error())
		return err
	}
	return nil
}
