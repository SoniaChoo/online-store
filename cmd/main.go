package main

import (
	myHttp "github.com/SoniaChoo/online-store/http"
	"net/http"
)

func main() {
	http.HandleFunc("/user/register", myHttp.RegisterHandler)
	http.HandleFunc("/user/login", myHttp.LoginHandler)
	http.HandleFunc("/user/retrieve/userid", myHttp.RetrieveIdHandler)
	http.HandleFunc("/rest/register", myHttp.RegisterHandlerRest)
	http.ListenAndServe(":12345", nil)
}
