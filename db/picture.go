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
	_, err = db.QueryContext(ctx, "insert into picture(types, pic_path, first_pic) values(?,?,?)",
		p.Types, p.PicPath, p.FirstPic)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}
