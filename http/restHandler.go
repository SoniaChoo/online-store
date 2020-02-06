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

const (
	BadJsonRest = "Restaurant register info wrong"
	RequestParameterMissingRest="userid should not be zero, phone/address/restaurantname/ should not be empty!"
	SuccessfullyRegisterRest = "Restaurant %s is successfully registered!"
	SuccessfullyShowdishesRest = "all dishes of restaurant %v are successfully showed as following: %v\n"
	)

func RegisterHandlerRest(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Restaurant request error!")
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
		fmt.Fprintf(w,RequestParameterMissingRest)
		return
	}

	//insert rest info into database
	if err = db.InsertRest(&rest); err != nil {
		log.Printf("creat rest failed, error is %s\n", err.Error())
		if strings.Contains(err.Error(), DuplicateError) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "This phone has been registered!")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Restaurant register failed!")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyRegisterRest, rest.Phone)
}

func ShowDishesHandlerRest(w http.ResponseWriter, r *http.Request) {
	reqbytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Restaurant request error!")
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
	if err != nil {
		log.Printf("show dishes failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "show dishes failed, can't show restaurant's rest's dishes, error is %s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyShowdishesRest, rest.RestId, dishes)
}

func RetrieveHandlerRest(w http.ResponseWriter, r *http.Request) {
	reqbytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Rest request error!")
		return
	}
	defer r.Body.Close() //remember to close network connection

	//unmarshal the byte data Rest info
	var rest model.Rest
	if err = json.Unmarshal(reqbytes, &rest); err != nil {
		log.Printf("Read rest info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonRest)
		return
	}

	//retrieve rest by restname in a fuzzy match way
	rests, err := db.RetrieveRest(&rest)
	if err != nil {
		log.Printf("retrieve rests failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "no result for this name, error is %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "all rests of contain %v are successfully showed as followed: %v\n", rest.RestName, rests)
}
