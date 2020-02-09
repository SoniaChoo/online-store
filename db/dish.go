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
	_, err = db.QueryContext(ctx, "insert into dish(rest_id, price, dish_name, description, stock) values(?,?,?,?,?)",
		d.RestId, d.Price, d.DishName, d.Description, d.Stock)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}

func ShowDishesDeatil(d *model.Dish) ([]*model.Dish, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	details := []*model.Dish{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from dish where dish_id = ?", d.DishId)
	if err != nil {
		log.Printf("record show dish detail by dishid %d with error %s\n", d.DishId, err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		temp := &model.Dish{}
		if err = row.Scan(&temp.DishId, &temp.RestId, &temp.Price, &temp.DishName, &temp.Description, &temp.Stock, &temp.Sales, &temp.Favorite, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("record loop show dish detail by dishid %d with error %s\n", d.DishId, err.Error())
			return nil, err
		}
		details = append(details, temp)
	}

	return details, nil
}

func UpdateDish(d *model.Dish) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	//in the SQL statement, "," can not be replaced by "and"
	_, err = db.QueryContext(ctx, "update dish set price = ?, dish_name = ?, description = ?, stock = ? where dish_id = ? ", d.Price, d.DishName, d.Description, d.Stock, d.DishId)
	if err != nil {
		log.Printf("record update dish with error %s\n", err.Error())
		return err
	}

	return nil
}

func SearchByDishNameDish(d *model.Dish) ([]*model.Dish, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	dishes := []*model.Dish{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from dish where dish_name like ?", "%"+d.DishName+"%")
	if err != nil {
		log.Printf("record search dish by dishname %s with error %s\n", d.DishName, err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		temp := &model.Dish{}
		if err = row.Scan(&temp.DishId, &temp.RestId, &temp.Price, &temp.DishName, &temp.Description, &temp.Stock, &temp.Sales, &temp.Favorite, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("record search dish by dishname %s in loop with error %s\n", d.DishName, err.Error())
			return nil, err
		}
		dishes = append(dishes, temp)
	}

	return dishes, nil
}
