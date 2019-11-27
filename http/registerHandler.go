package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SoniaChoo/online-store/service"

	"github.com/SoniaChoo/online-store/db"
	"github.com/SoniaChoo/online-store/model"
)

// RegisterHandler is the function for user register,
// parameters: email, passwd.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
	var user model.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("Read user info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User register info wrong!")
		return
	}

	// check user variable

	// insert user info into database
	if err = db.InsertUser(user); err != nil {
		log.Printf("create user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User register failed!")
		return
	}

	// write response to client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s successfully registered!", user.Email)

}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
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
	var changeuser model.ChangeUser
	if err = json.Unmarshal(reqBytes, &changeuser); err != nil {
		log.Printf("Read user info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "changeuser update info wrong!")
		return
	}

	// check user variable

	// insert user info into database
	oldU := model.User{
		changeuser.Email,
		changeuser.Passwd,
	}
	newU := model.User{
		changeuser.NewEmail,
		changeuser.NewPasswd,
	}
	if err = db.UpdateUser(oldU, newU); err != nil {
		log.Printf("update user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User update failed!")
		return
	}

	// write response to client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s successfully updated!", oldU.Email)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
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
	var user model.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("Read user info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "deleteuser info wrong!")
		return
	}

	// check user variable

	// delete user info into database

	if err = db.DeleteUser(user); err != nil {
		log.Printf("delete user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User delete failed!")
		return
	}

	// write response to client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s successfully deleted!", user.Email)

}

// TODO
func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
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
	var user1 model.User
	if err = json.Unmarshal(reqBytes, &user1); err != nil {
		log.Printf("Read user info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "retrive user info wrong!")
		return
	}

	// check user variable

	// delete user info into database
	var user model.User
	if user, err = db.RetrieveUser(user1.Email); err != nil {
		log.Printf("retrive user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User retrive failed!")
		return
	}

	// write response to client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s successfully retrive!", user.Email)

}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	nowHour, nowMinute, goWorkRes1, goWorkRes2, backHomeRes1, backHomeRes2 := service.GetTime()

	// write response to client
	w.WriteHeader(http.StatusOK)
	outputF := "Nowtime: %02d:%02d\n\n\n\nWORK, from Nakai to Roppongi, next: %d minutes, next next %d minutes.\n\n\n\nHOME, from Roppongi to Nakai, next: %d minutes, next next %d minutes.\n\n\n\n"
	fmt.Fprintf(w, outputF, nowHour, nowMinute, goWorkRes1, goWorkRes2, backHomeRes1, backHomeRes2)
}
