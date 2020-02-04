package http

import (
	"encoding/json"
	"fmt"
	"github.com/SoniaChoo/online-store/db"
	"github.com/SoniaChoo/online-store/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const BadJsonRest = "Rest register info wrong"

func RegisterHandlerRest(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Rest request error!")
		return
	}
	defer r.Body.Close() //remember to close network connection

	//unmarshal the byte data into Rest info
	var rest model.Rest
	if err = json.Unmarshal(reqBytes, &rest); err != nil {
		log.Printf("Read rest info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonRest)
		return
	}

	//check rest variable
	if rest.UserId == 0 || rest.Phone == "" || rest.Address == "" || rest.RestName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "userid should not be zero, phone/address/restname/ should not be empty!")
		return
	}

	//insert rest info into database
	if err = db.InsertRest(&rest); err != nil {
		log.Printf("creat rest failed, error is %s\n", err.Error())
		if strings.Contains(err.Error(), DuplicateError) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "This phone has been registered!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Rest register failed!")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Rest %v is successfully registered!", rest)
}

func ShowDishesHandlerRest(w http.ResponseWriter, r *http.Request) {
	reqbytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Rest Request error!")
		return
	}
	defer r.Body.Close() //remember to closer network connrction

	//unmarshal the byte data into Rest info
	var rest model.Rest
	if err = json.Unmarshal(reqbytes, &rest); err != nil {
		log.Printf("Read rest info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonRest)
		return
	}

	//search all dishes belong to the rest
	dishes, err := db.ShowDishesRest(&rest)
	if  err != nil {
		log.Printf("show dishes failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "show dishes failed, can't show this rest's dishes, error is %s\n", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "all dishes of rest %v are successfully showed as following: %v\n", rest.RestId, dishes)
}

