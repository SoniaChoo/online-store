package db

import (
	"context"
	"github.com/SoniaChoo/online-store/model"
	"github.com/pkg/errors"
	"log"
	"time"
)

const NotMatchError = "sql: no rows in result set"

// InsertUser is to insert a user info into database
func InsertUser(u *model.User) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return err
	}

	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into user(phone, nickname, password) values(?,?,?)",
		u.Phone, u.Nickname, u.Password)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}

func LoginUser(u *model.User) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	row, err := db.QueryContext(ctx, "select * from user where phone = ? and password = ? ", u.Phone, u.Password)
	if err != nil {
		log.Printf("record login with error %s\n", err.Error())
		return err
	}
	defer row.Close()

	if !row.Next() {
		log.Printf(NotMatchError)
		return errors.New(NotMatchError)
	}

	return nil
}

func RetrieveUserId(u *model.User) ([]*model.User, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	users := []*model.User{}
	row, err := db.QueryContext(ctx, "select * from user where user_id = ?", u.UserId)
	if err != nil {
		log.Printf("record retrieve with error %s\n", err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		temp := &model.User{}
		if err = row.Scan(&temp.UserId, &temp.Phone, &temp.Nickname, &temp.Password, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("retrieving record loop with error %s\n", err.Error())
			return nil, err
		}
		users = append(users, temp)
	}

	return users, nil
}

func RetrieveUserPhone(u *model.User) ([]*model.User, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
		return nil, err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	users := []*model.User{}
	row, err := db.QueryContext(ctx, "select * from user where phone = ?", u.Phone)
	if err != nil {
		log.Printf("record retrieve with error %s\n", err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		temp := &model.User{}
		if err = row.Scan(&temp.UserId, &temp.Phone, &temp.Nickname, &temp.Password, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("retrieving record loop with error %s\n", err.Error())
			return nil, err
		}
		users = append(users, temp)
	}

	return users, nil
}

func RetrieveUserNickname(u *model.User) ([]*model.User, error) {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err) //为什么err,err.Error
		return nil, err
	}

	//start to excute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	users := []*model.User{}
	row, err := db.QueryContext(ctx, "select * from user where nickname like ?", "%"+u.Nickname+"%")
	if err != nil {
		log.Printf("record retrieve with error %s\n", err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		temp := &model.User{}
		if err = row.Scan(&temp.UserId, &temp.Phone, &temp.Nickname, &temp.Password, &temp.CreatTime, &temp.UpdateTime); err != nil {
			log.Printf("retrieving record loop with error %s\n", err.Error())
			return nil, err
		}
		users = append(users, temp)
	}

	return users, nil
}
