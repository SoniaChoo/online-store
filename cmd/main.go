package main

import (
	"net/http"

	myHttp "github.com/SoniaChoo/online-store/http"
)

func main() {
	http.HandleFunc("/", myHttp.RegisterHandler)
	http.ListenAndServe(":12345", nil)
}
