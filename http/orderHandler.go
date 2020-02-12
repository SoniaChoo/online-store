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
	BadJsonShowCart      = "Show order info wrong"
	SuccessfullyShowCart = "the cart of orderid = %d is successfully as following: %v"
)

func ShowCartHandlerOrder(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read requset error! Error is %\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Order request error!")
		return
	}

	//unmarshal the byte into order info
	var order model.Orders
	if err = json.Unmarshal(reqBytes, &order); err != nil {
		log.Printf("Read order info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonShowCart)
		return
	}

	//select by userid in order to get orderid
	orderids, err := db.GetOrderIdInTableOrder(&order)
	if err != nil {
		log.Printf("select by userid in order to get order_id with error, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "select by userid = %d in table order failed")
		return
	}

	//judge if len(details) = 1
	if len(orderids) == 0 || len(orderids) > 1 {
		log.Printf("length of orderids should be 1")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "the result of this orderids should be unique, but it's length = %d", len(orderids))
		return
	}
	orderid := orderids[0].OrderId

	//select by orderid in order_detail to show cart
	carts, err := db.ShowCartOrder(orderid)
	if err != nil {
		log.Printf("select by orderid in order_detail to show cart with error, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "select by orderid = %d in order_detail failed", orderid)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyShowCart, orderid, carts)
}