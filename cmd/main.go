package main

import (
	"net/http"

	myHttp "github.com/SoniaChoo/online-store/http"
)

func main() {
	http.HandleFunc("/user/register", myHttp.RegisterHandler)
	http.HandleFunc("/", myHttp.TimeHandler)
	http.ListenAndServe(":80", nil)
}
