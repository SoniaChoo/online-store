package db

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"
)

const NotMatchError = "sql: no rows in result set"

// InsertUser is to insert a user info into database
func InsertUser(u *model.User) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
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
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	row, err := db.QueryContext(ctx, "select * from user where phone = ? and password = ? ", u.Phone, u.Password)
	if err != nil {
		log.Printf("record login with error %s\n", err.Error())
		return err
	}
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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	//store the result by slice
	user := []*model.User{}
	row, err := db.QueryContext(ctx, "select * from user where user_id = ?", u.UserId)
	if err != nil {
		log.Printf("record retrieve with error %s\n", err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		temp := &model.User{}
		if err = row.Scan(&temp.UserId, &temp.Phone, &temp.Nickname, &temp.Password, &temp.CreatTime); err != nil {
			log.Printf("retrieving record loop with error %s\n", err.Error())
			return nil, err
		}
		user = append(user, temp)
	}

	return user, nil
}

// TODO: implement update function.
//func UpdateUser(user, newUser model.User) error {
//	db, err := DBFactory()
//	if err != nil {
//
//	}
//	// start to execute SQL query
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
//	defer cancel()
//	_, err = db.QueryContext(ctx, "update user set email = ?, passwd = ? where email = ? and passwd = ?",
//		newUser.Email, newUser.Passwd, user.Email, user.Passwd)
//	if err != nil {
//		log.Printf("record updating with error %s\n", err.Error())
//		return err
//	}
//
//	return nil
//}
//
////delete user info
//func DeleteUser(user model.User) error {
//	db, err := DBFactory()
//	if err != nil {
//
//	}
//
//	// start to execute SQL query
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
//	defer cancel()
//	_, err = db.QueryContext(ctx, "delete from user where email = ? and passwd = ? ",
//		user.Email, user.Passwd)
//	if err != nil {
//		log.Printf("record deleting with error %s\n", err.Error())
//		return err
//	}
//
//	return nil
//}
//
//func RetrieveUser(email string) (model.User, error) {
//	db, err := DBFactory()
//	if err != nil {
//
//	}
//
//	// start to execute SQL query
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
//	defer cancel()
//	rows, err := db.QueryContext(ctx, "select * from user where email = ? limit 1", email)
//	if err != nil {
//		log.Printf("record retrieve with error %s\n", err.Error())
//		return model.User{}, err
//	}
//	defer rows.Close()
//
//	var user model.User
//	for rows.Next() {
//		var (
//			vid     int
//			vemail  string
//			vpasswd string
//			vc      []uint8
//			vu      []uint8
//		)
//		err = rows.Scan(&vid, &vemail, &vpasswd, &vc, &vu)
//		if err != nil {
//			log.Printf("retriving record loop with error %s\n", err.Error())
//			return model.User{}, err
//		}
//		user = model.User{
//			vemail,
//			vpasswd,
//		}
//	}
//
//	return user, nil
//}
