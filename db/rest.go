package db

import (
	"context"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"
)

// InsertRest is to insert a rest info into database
func InsertRest(r *model.Rest) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}

	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into rest(user_id, phone, rest_name, address) values(?,?,?,?)",
		r.UserId, r.Phone, r.RestName, r.Address)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}

func ShowDishesRest(r *model.Rest) ([]*model.Dish, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	dishes := []*model.Dish{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from dish where rest_id = ?", r.RestId)
	if err != nil {
		log.Printf("record show dished of rest with error %s\n", err.Error())
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		temp := &model.Dish{}
		if err = row.Scan(&temp.DishId, &temp.RestId, &temp.Price, &temp.DishName, &temp.Description, &temp.Stock, &temp.Sales, &temp.Favorite, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("doing show dishes record loop with error %s\n", err.Error())
			return nil, err
		}
		dishes = append(dishes, temp)
	}

	return dishes, nil
}

func RetrieveRest(r *model.Rest) ([]*model.Rest, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	rests := []*model.Rest{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from rest where rest_name like ?", "%"+r.RestName+"%")
	if err != nil {
		log.Printf("record retrieve rest by restname with error %s\n", err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		temp := &model.Rest{}
		if err = row.Scan(&temp.RestId, &temp.UserId, &temp.Phone, &temp.RestName, &temp.Address, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("record retrieving rest loop with error %s\n", err.Error())
			return nil, err
		}
		rests = append(rests, temp)
	}

	return rests, nil
}
