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
	BadJsonShowCart                  = "Show order info wrong"
	SuccessfullyShowCart             = "the cart of orderid = %d is successfully as following: %v"
	BadJsonAddToCart                 = "Add to cart info wrong"
	SuccessfullyAddToCart            = "the dish dishid = %d is successfully added into cart"
	BadJsonDeleteDishInCart          = "Delete dish in cart info wrong"
	SuccessfullyDeleteOrUpdateInCart = "dish = %d is successfully delete in cart"
)

func ShowCartHandlerOrder(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read requset error! Error is %s\n", err.Error())
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
	orderids, err := db.GetOrderIdInTableOrder(order.UserId)
	if err != nil {
		log.Printf("select by userid in order to get order_id with error, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "select by userid = %d in table order failed", order.UserId)
		return
	}

	//judge if len(details) = 1, must be 1, because every userid always has a cart since register.if not one, code has problem,it's different from Showdishdetail
	if len(orderids) != 1 {
		log.Printf("orderids should be existed, and length of it should be 1")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "the result of this orderids should be exist and unique, but it's length = %d", len(orderids))
		return
	}
	orderId := orderids[0].OrderId

	//select by orderid in order_detail to show cart
	carts, err := db.ShowCartOrder(orderId)
	if err != nil {
		log.Printf("select by orderid in order_detail to show cart with error, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "select by orderid = %d in order_detail failed", orderId)
		return
	}

	//loop the cart to get the total_price in table order and update
	totalPrice := float64(0)
	for i := 0; i < len(carts); i++ {
		totalPrice += carts[i].Price * float64(carts[i].Number)
	}

	//update total_price by order_id in table orders
	if err = db.UpdateTotalPriceInOrder(totalPrice, orderId); err != nil {
		log.Printf("update total_price by order_id in table orders with error, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "update total_price by order_id in table orders failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyShowCart, orderId, carts)
}

func AddToCartHandlerOrder(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Order request error!")
		return
	}

	//unmarshal the byte into order info
	//var order model.Orders
	type addToCart struct {
		Userid int     `json:"userid"`
		Dishid int     `json:"dishid"`
		Restid int     `json:"restid"`
		Price  float64 `json:"price"`
		Number int     `json:"number"`
	}
	var atc = addToCart{}
	if err = json.Unmarshal(reqBytes, &atc); err != nil {
		log.Printf("Read order info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonAddToCart)
		return
	}

	//select by userid in order to get orderid
	orderids, err := db.GetOrderIdInTableOrder(atc.Userid)
	if err != nil {
		log.Printf("select by userid in order to get order_id with error, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "select by userid = %d in table order failed", atc.Userid)
		return
	}

	//judge if len(orderids)=1, must be 1, because every userid always has a cart since register, if not one, code has bug
	if len(orderids) != 1 {
		log.Printf("orderids should be existed, and length of it should be 1")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "the result of this orderids should br exist and unique, but it's length = %d", len(orderids))
		return
	}
	orderId := orderids[0].OrderId
	detail := model.Order_detail{
		DishId:  atc.Dishid,
		RestId:  atc.Restid,
		OrderId: orderId,
		Price:   atc.Price,
		Number:  atc.Number,
	}

	// insert or update
	if err = db.AddToCartOrder(&detail); err != nil {
		log.Printf("add dish to cart with error, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "dishid = %d add dish to cart failed", detail.DishId)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyAddToCart, detail.DishId)
}

func DeleteDishInCartHandlerOrder(w http.ResponseWriter, r *http.Request) {
	//get the request info
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Order request error!")
		return
	}

	//unmarshal the byte into order info
	var detail model.Order_detail
	if err = json.Unmarshal(reqBytes, &detail); err != nil {
		log.Printf("Read order info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonDeleteDishInCart)
		return
	}

	if err = db.DeleteOrUpdateDishInCart(&detail); err != nil {
		log.Printf("delete or update dish in cart with error, error is %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "detailid = %d delete or update dish in cart failed", detail.DetailId)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, SuccessfullyDeleteOrUpdateInCart, detail.DishId)
}
