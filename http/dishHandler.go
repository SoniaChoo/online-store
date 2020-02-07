package http

import (
	"encoding/json"
	"fmt"
	"github.com/SoniaChoo/online-store/db"
	"github.com/SoniaChoo/online-store/model"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	SuccessfullyShowdishDetail  = "%v dish's detail is successfully showed ad following: %v\n"
	BadJsonDish                 = "Dish show detail info wrong"
	RequestParameterMissingDish = "RestId, DishPrice, DishStock should not be zero, DishName, Description should not be empty"
	SuccessfullyAddDish         = "Dish %s is successfully added!"
)

func DetailHandlerDish(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Dish request error!")
		return
	}
	defer r.Body.Close()

	//unmarshal the byte data into Rest info
	var dish model.Dish
	if err = json.Unmarshal(reqBytes, &dish); err != nil {
		log.Printf("Read dish info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonDish)
		return
	}

	//retrieve dish detail from database
	details, err := db.ShowDishesDeatil(&dish)
	if err != nil {
		log.Printf("show dish detail failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "this dish's detail can not be showed, error is %s\n", err.Error())
		return
	}

	//judge if len(details) = 1
	if len(details) != 1 {
		log.Printf("len(details) should be 1")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "the result of this dish should be unique, but it's not, error is %s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyShowdishDetail, dish.DishId, details[0])
}

func AddHandlerDish(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Dish request error")
		return
	}
	defer r.Body.Close()

	//unmarshal the byte data into Rest info
	var dish model.Dish
	if err = json.Unmarshal(reqBytes, &dish); err != nil {
		log.Printf("Read dish info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonDish)
		return
	}

	//check rest variable
	if dish.RestId == 0 || dish.Price == 0 || dish.DishName == "" || dish.Description == "" || dish.Stock == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, RequestParameterMissingDish)
		return
	}

	//insert dish info into database
	if err = db.InsertDish(&dish); err != nil {
		log.Printf("add dish failed, error is %s\n", err.Error())
		//if strings.Contains(err.Error(), DuplicateError) {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	fmt.Fprintf(w, "This dish had been added")
		//	return
		//}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Dish added failed!")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyAddDish, dish.DishName)
}
