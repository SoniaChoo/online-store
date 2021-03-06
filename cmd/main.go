package main

import (
	myHttp "github.com/SoniaChoo/online-store/http"
	"net/http"
)

func main() {
	http.HandleFunc("/user/register", myHttp.RegisterHandler)
	http.HandleFunc("/user/login", myHttp.LoginHandler)
	http.HandleFunc("/user/retrieve/userid", myHttp.RetrieveIdHandler)
	http.HandleFunc("/user/retrieve/phone", myHttp.RetrievePhoneHandler)
	http.HandleFunc("/user/retrieve/nickname", myHttp.RetrieveNicknameHandler)
	http.HandleFunc("/rest/register", myHttp.RegisterHandlerRest)
	http.HandleFunc("/rest/showdishes", myHttp.ShowDishesHandlerRest)
	http.HandleFunc("/rest/retrieve", myHttp.RetrieveHandlerRest)
	http.HandleFunc("/dish/detail", myHttp.DetailHandlerDish)
	http.HandleFunc("/dish/add", myHttp.AddHandlerDish)
	http.HandleFunc("/dish/update", myHttp.UpdateHandlerDish)
	http.HandleFunc("/dish/search/name", myHttp.SearchByDishNameHandlerDish)
	http.HandleFunc("/dish/search/description", myHttp.SearchByDescriptionHandlerDish)
	http.HandleFunc("/order/showcart", myHttp.ShowCartHandlerOrder)
	http.HandleFunc("/order/addtocart", myHttp.AddToCartHandlerOrder)
	http.HandleFunc("/order/deletedishincart", myHttp.DeleteDishInCartHandlerOrder)
	http.ListenAndServe(":12345", nil)
}
