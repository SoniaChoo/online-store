package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SoniaChoo/online-store/config"

	"github.com/SoniaChoo/online-store/model"

	_ "github.com/go-sql-driver/mysql"
)

// InsertUserInfo is to insert a userinfo into database
func InsertUserInfo(u model.UserInfo) error {//zhuzhu:修改了这里的函数名字，更新了参数的类型
	// read environment config
	env, err := config.ReadFromEnv()
	if err != nil {
		log.Printf("read environment variable fail, error is %s\n", err.Error())
		return err
	}

	// build database config string from env
	var dataSourName strings.Builder
	fmt.Fprint(&dataSourName, env.DBUser)
	fmt.Fprint(&dataSourName, ":")
	fmt.Fprint(&dataSourName, env.DBPass)
	fmt.Fprint(&dataSourName, "@tcp(")
	fmt.Fprint(&dataSourName, env.DBHost)
	fmt.Fprint(&dataSourName, ":")
	fmt.Fprint(&dataSourName, env.DBPort)
	fmt.Fprint(&dataSourName, ")/")
	fmt.Fprint(&dataSourName, env.DBName)
	//log.Printf("database configuration string is %s\n", dataSourName.String())

	// connect database
	db, err := sql.Open("mysql", dataSourName.String())
	if err != nil {
		log.Printf("open database failed: %s\n", err.Error())
		return err
	}
	db.SetConnMaxLifetime(0) // to set connections reuse forever
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(0)
	log.Println("database connected.")

	// start to execute SQL query
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = db.QueryContext(ctx, "insert into userinfo(name, age, gender, address) values(?,?,?,?)",//zhuhzu:这里面的insert into table_name也更新了
		u.Name, u.Age, u.Gender, u.Address)//zhuzhu:更新了字段名,上一行也更新了
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}


//func Insert(user model.User) error {}