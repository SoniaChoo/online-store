package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/SoniaChoo/online-store/config"
	_ "github.com/go-sql-driver/mysql"
)

var globalDB *sql.DB
var once sync.Once

//将连接database的过程写为一个函数.
func connectDB() (*sql.DB, error) {
	// read environment config
	env, err := config.ReadFromEnv()
	if err != nil {
		log.Printf("read environment variable fail, error is %s\n", err.Error())
		return nil, err
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
		return nil, err
	}
	db.SetConnMaxLifetime(0) // to set connections reuse forever
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(0)
	log.Println("database connected.")

	return db, nil
}

func DBFactory() (*sql.DB, error) {
	once.Do(func() {
		var err error
		if globalDB, err = connectDB(); err != nil {
			log.Printf("open database failed: %s\n", err.Error())
		}
	})

	return globalDB, nil
}
