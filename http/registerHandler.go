package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
		log.Fatalf("Read request error! Error is %s\n", err.Error())
	}
	defer r.Body.Close() // remember to close network connection

	// unmarshal the byte data into User info
	var user model.User
	err = json.Unmarshal(reqBytes, &user)
	if err != nil {
		log.Fatalf("Read user info error! Error is %s\n", err.Error())
	}

	// TODO: insert user info into database
	time.Sleep(1 * time.Second) //simulate database operation

	// write response to client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s successfully registered!", user.GetEmail())
}
