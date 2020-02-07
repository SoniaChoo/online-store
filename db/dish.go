package db

import (
	"context"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"
)

// InsertDish is to insert a dish info into database
func InsertDish(d *model.Dish) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}

	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into dish(rest_id, price, dish_name, stock, sales) values(?,?,?,?,?)",
		d.RestId, d.Price, d.DishName, d.Stock, d.Sales)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}

func ShowDishesDeatil(d *model.Dish) (*model.Dish, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	detail := &model.Dish{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from dish where dish_id = ?", d.DishId)
	if err != nil {
		log.Printf("record show dishes edtail with error %s\n", err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		temp := model.Dish{}
		if err = row.Scan(&temp.DishId, &temp.RestId, &temp.Price, &temp.DishName, &temp.Description, &temp.Stock, &temp.Sales, &temp.Favorite, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("record show dish detail loop with error %s\n", err.Error())
			return nil, err
		}
		detail = &temp
	}

	return detail, nil
}
