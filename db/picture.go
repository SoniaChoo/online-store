package db

import (
	"context"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"
)

// InsertPicture is to insert a picture info into database
func InsertPicture(p model.Picture) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}
	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into user(pic_id, dish_id, user_id, rest_id, pic_path, first_pic, creat_time) values(?,?,?,?,?,?,?)",
		p.PicId, p.DishId, p.UserId, p.RestId, p.PicPath, p.FirstPic, p.CreatTime)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}