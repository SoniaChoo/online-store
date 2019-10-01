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

// TODO: implement update function.
func UpdateUser(user, newUser model.User) error {
// func UpdateUser(u model.User) error {
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
		_, err = db.QueryContext(ctx, "delete from user where email = ? and passwd = ? ",
		 user.Email, user.Passwd)
		if err != nil {
			log.Printf("record deleting with error %s\n", err.Error())
			return err
		}
	
		return nil
	}

func RetrieveUser(email string) (model.User,error) {
		// read environment config
		env, err := config.ReadFromEnv()
		if err != nil {
			log.Printf("read environment variable fail, error is %s\n", err.Error())
			return  model.User{}, err
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
			return model.User{}, err
		}
		db.SetConnMaxLifetime(0) // to set connections reuse forever
		db.SetMaxIdleConns(0)
		db.SetMaxOpenConns(0)
		log.Println("database connected.")
	
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
				vid int
				vemail   string
				vpasswd string
				vc []uint8
				vu []uint8
			)
			err = rows.Scan(&vid, &vemail, &vpasswd, &vc, &vu)
			if err != nil{
				log.Printf("retriving record loop with error %s\n", err.Error())
				return model.User{}, err
			}
			user = model.User {
				vemail,
				vpasswd,
			}
		}
	
		return user, nil
	}	