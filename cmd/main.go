package main

import (
	"log"
	"net/http"

	myHttp "github.com/SoniaChoo/online-store/http"
)

func main() {
	//http.HandleFunc("/user/register", myHttp.RegisterHandler)
	http.HandleFunc("/", myHttp.TimeHandler)
	log.Fatal(http.ListenAndServe(":443", nil))
}
