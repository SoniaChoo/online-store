package db

import (
	"context"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"
)

// InsertUser is to insert a user info into database
func InsertUser(u model.User) error {
	db, err := DBFactory()
	if err != nil {
		log.Printf("error connect database, %v\n", err)
	}
	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into user(user_id, phone, nickname,password,creat_time) values(?,?,?,?,?)",
		u.UserId, u.Phone,u.Nickname,u.Password, u.CreatTime)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}

// TODO: implement update function.
func UpdateUser(user, newUser model.User) error {
	db, err := DBFactory()
	if err != nil {

	}
	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "update user set email = ?, passwd = ? where email = ? and passwd = ?",
		newUser.Email, newUser.Passwd, user.Email, user.Passwd)
	if err != nil {
		log.Printf("record updating with error %s\n", err.Error())
		return err
	}

	return nil
}

//delete user info
func DeleteUser(user model.User) error {
	db, err := DBFactory()
	if err != nil {

	}

	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "delete from user where email = ? and passwd = ? ",
		user.Email, user.Passwd)
	if err != nil {
		log.Printf("record deleting with error %s\n", err.Error())
		return err
	}

	return nil
}

func RetrieveUser(email string) (model.User, error) {
	db, err := DBFactory()
	if err != nil {

	}

	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	rows, err := db.QueryContext(ctx, "select * from user where email = ? limit 1", email)
	if err != nil {
		log.Printf("record retrieve with error %s\n", err.Error())
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		var (
			vid     int
			vemail  string
			vpasswd string
			vc      []uint8
			vu      []uint8
		)
		err = rows.Scan(&vid, &vemail, &vpasswd, &vc, &vu)
		if err != nil {
			log.Printf("retriving record loop with error %s\n", err.Error())
			return model.User{}, err
		}
		user = model.User{
			vemail,
			vpasswd,
		}
	}

	return user, nil
}
