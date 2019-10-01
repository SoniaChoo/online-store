package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SoniaChoo/online-store/db"
	"github.com/SoniaChoo/online-store/model"
)

// userinfoHandler is the function for inputing userinfo,
// parameters: name,age, gender, address.
func UserinfoHandler(w http.ResponseWriter, r *http.Request) {//zhuzhu:更新了函数名
	// check the request info, not needed, just for understanding.
	log.Println("Method:", r.Method)
	log.Println("URL:", r.URL, "URL.Path", r.URL.Path)
	log.Println("RemoteAddress", r.RemoteAddr)
	log.Println("UserAgent", r.UserAgent())
	log.Println("Header", r.Header)
	log.Println("Cookies", r.Cookies())

	// get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User request error!")
		return
	}
	defer r.Body.Close() // remember to close network connection

	// unmarshal the byte data into User info
	var user model.UserInfo//zhuzhu:更新了user的类型
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("Read user info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Inputing UserInfo wrong!")
		return
	}

	// insert user info into database
	if err = db.InsertUserInfo(user); err != nil {//zhuzhu:更新了函数的名字
		log.Printf("create user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User register failed!")
		return
	}

	// write response to client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s successfully registered!", user.Name)//zhuzhu:这里也要改
}
