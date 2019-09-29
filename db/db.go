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

// InsertUser is to insert a user info into database
func InsertUser(u model.User) error {
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
	_, err = db.QueryContext(ctx, "insert into user(email, passwd) values(?,?)",
		u.Email, u.Passwd)
	if err != nil {
		log.Printf("record inserting with error %s\n", err.Error())
		return err
	}

	return nil
}
