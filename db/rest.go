package db

import (
	"context"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"
)

// InsertRest is to insert a rest info into database
func InsertRest(r model.Rest) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}
	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into user(rest_id, user_id, phone, rest_name, address, creat_time) values(?,?,?,?,?,?)",
		r.RestId, r.UserId, r.Phone, r.RestName, r.Address, r.CreatTime)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}