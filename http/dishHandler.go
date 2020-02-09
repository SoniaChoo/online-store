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
	SuccessfullyShowdishDetail          = "%v dish's detail is successfully showed ad following: %v\n"
	BadJsonDishDetail                   = "Dish show detail info wrong"
	RequestParameterMissingDish         = "RestId, DishPrice, DishStock should not be zero, DishName, Description should not be empty"
	BadJsonDishAdd                      = "Dish add info wrong"
	SuccessfullyAddDish                 = "Dish %s is successfully added!"
	BadJsonDishUpdate                   = "Dish update info wrong"
	SuccessfullyUpdateDish              = "Dish %d is successfully updated!"
	RequestUpdateNotAvailable           = "after updated, price/dishname/description/stock should be zero or empty"
	BadJsonDishSearch                   = "Dish search info wrong "
	SuccessfullySearchByDishNameDish    = "dishname like %s is successfully searched as following: %v\n"
	SuccessfullySearchByDescriptionDish = "description like %s is successfully searched as following: %v\n"
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
		fmt.Fprintf(w, BadJsonDishDetail)
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
		fmt.Fprintf(w, BadJsonDishAdd)
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

func UpdateHandlerDish(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Dish request error!")
		return
	}
	defer r.Body.Close()

	//unmarshal the byte data into dish info
	var dish model.Dish
	if err = json.Unmarshal(reqBytes, &dish); err != nil {
		log.Printf("Read dish info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonDishUpdate)
		return
	}

	//check avariable
	if dish.Price == 0 || dish.DishName == "" || dish.Description == "" || dish.Stock == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, RequestUpdateNotAvailable)
		return
	}

	//update dish detail
	if err = db.UpdateDish(&dish); err != nil {
		log.Printf("update dish failed, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Dish updated failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyUpdateDish, dish.DishId)
}

func SearchByDishNameHandlerDish(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Dish request error!")
		return
	}
	defer r.Body.Close()

	//unmarshal the byte into dish data
	var dish model.Dish
	if err = json.Unmarshal(reqBytes, &dish); err != nil {
		log.Printf("Read dish info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonDishSearch)
		return
	}

	//search by dishname in database
	dishes, err := db.SearchByDishNameDish(&dish)
	if err != nil {
		log.Printf("search dishs by dishname %s failed, error is %s\n", dish.DishName, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "search dishs by dish name %s failed, error is %s\n", dish.DishName, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullySearchByDishNameDish, dish.DishName, dishes)
}

func SearchByDescriptionHandlerDish(w http.ResponseWriter, r *http.Request) {
	//read the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Dish request error!")
		return
	}
	defer r.Body.Close()

	//unarshal the bytes into dish data
	var dish model.Dish
	if err = json.Unmarshal(reqBytes, &dish); err != nil {
		log.Printf("Read dish info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonDishSearch)
		return
	}

	//search by description in database
	dishs, err := db.SearchByDescriptionDish(&dish)
	if err != nil {
		log.Printf("search dishs by description %s failed, error is %s\n", dish.Description, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "searh dishs by description %s failed, error is %s\n", dish.Description, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullySearchByDescriptionDish, dish.Description, dishs)
}
