package db

import (
	"context"
	"log"
	"time"

	"github.com/SoniaChoo/online-store/model"

	_ "github.com/go-sql-driver/mysql"
)

// InsertUserInfo is to insert a userinfo into database
func InsertUserInfo(u model.UserInfo) error { //zhuzhu:修改了这里的函数名字，更新了参数的类型
	db, err := DBFactory()
	if err != nil {

	}

	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into userinfo(name, age, gender, address) values(?,?,?,?)", //zhuhzu:这里面的insert into table_name也更新了
		u.Name, u.Age, u.Gender, u.Address) //zhuzhu:更新了字段名,上一行也更新了
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}

//func Insert(user model.User) error {}
