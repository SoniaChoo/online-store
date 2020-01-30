package main

import (
	"net/http"

	myHttp "github.com/SoniaChoo/online-store/http"
)

func main() {
	http.HandleFunc("/user/register", myHttp.RegisterHandler)
	http.HandleFunc("/user/login", myHttp.LoginHandler)
	http.HandleFunc("/user/retrieve/userid", myHttp.RetrieveIdHandler)
	http.ListenAndServe(":12345", nil)
}
