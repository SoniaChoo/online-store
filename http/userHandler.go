package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/SoniaChoo/online-store/db"
	"github.com/SoniaChoo/online-store/model"
)

const (
	DuplicateError                  = "Error 1062"
	BadJsonRegister                 = "User register info wrong!"
	RequestRegisterParameterMissing = "phone/nickname/password should not be empty!"
	SuccessfullyRegister            = "User %s successfully registered!"
	BadJsonLogin                    = "User Login fail!"
	RequestLoginParameterMissing    = "phone/password should not be empty!"
	SuccessfullyLogin               = "User %s successfully login!"
	BadJsonRetrieve                 = "User retrieve wrong"
	SuccessfullyRetrieveId          = "User %v is successfully retrieved!"
)

// RegisterHandler is the function for user register,
// parameters: phone, nickname, password.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, BadJsonRegister)
		return
	}

	// check user variable
	if user.Phone == "" || user.Password == "" || user.Nickname == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, RequestRegisterParameterMissing)
		return
	}

	// insert user info into database
	if err = db.InsertUser(&user); err != nil {
		log.Printf("create user failed, error is %s\n", err.Error())
		if strings.Contains(err.Error(), DuplicateError) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "This phone number has been registered!")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User register failed!")
		return
	}

	//get userid
	users, err := db.RetrieveUserPhone(&user)
	if err != nil {
		log.Printf("creat user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "retrieve user by phone %s failed!", user.Phone)
		return
	}

	//judge if the length of users equal to 1, and get userid
	if len(users) == 0 || len(users) > 1 {
		log.Printf("len(users) should be 1")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "the result of this user should be unique, but it's %d", len(users))
		return
	}
	userid := users[0].UserId

	//use userid to new a cart
	if err = db.NewCart(userid); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "new a cart failed, error is %s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyRegister, user.Phone)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User request error!")
		return
	}
	defer r.Body.Close()

	var user model.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("Read user info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonLogin)
		return
	}

	if user.Phone == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, RequestLoginParameterMissing)
		return
	}

	if err = db.LoginUser(&user); err != nil {
		log.Printf("login user failed, error is %s\n", err.Error())
		if strings.Contains(err.Error(), db.NotMatchError) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "phone/password has not been matched")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User login failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyLogin, user.Phone)
}

func RetrieveIdHandler(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User request error!")
		return
	}
	defer r.Body.Close() //remember to close network connection

	//unmarshal the byte data into User info
	var user model.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("Read user info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonRetrieve)
		return
	}

	// retrieve user from database
	users, err := db.RetrieveUserId(&user)
	if err != nil {
		log.Printf("retrieve user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User retrieve failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyRetrieveId, users)
}

func RetrievePhoneHandler(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User request error!")
		return
	}
	defer r.Body.Close() //ember to close network connection

	//unmarshal the byte data into User info
	var user model.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("Read user info error!, Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User retrieve wrong")
		return
	}

	//retrieve user from database
	users, err := db.RetrieveUserPhone(&user)
	if err != nil {
		log.Printf("retrieve user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User retrieve failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %v is successfully retrieved!", users)
}

func RetrieveNicknameHandler(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is  %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User request error!")
		return
	}
	defer r.Body.Close()

	//unmarshal the byte data into User info
	var user model.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		log.Printf("retrieve user failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonRetrieve)
		return
	}

	//retrieve user from database
	users, err := db.RetrieveUserNickname(&user)
	if err != nil {
		log.Printf("retrieve user failed")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "User retrieve failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %v is successfully retrieved!", users)
}
