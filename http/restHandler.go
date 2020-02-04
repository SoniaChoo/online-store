package http

import (
	"encoding/json"
	"fmt"
	"github.com/SoniaChoo/online-store/db"
	"github.com/SoniaChoo/online-store/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const BadJsonRest = "Rest register info wrong"

func RegisterHandlerRest(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Rest request error!")
		return
	}
	defer r.Body.Close() //remember to close network connection

	//unmarshal the byte data into Rest info
	var rest model.Rest
	if err = json.Unmarshal(reqBytes, &rest); err != nil {
		log.Printf("Read rest info error! Error is %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, BadJsonRest)
		return
	}

	//check rest variable
	if rest.Phone == "" || rest.Address == "" || rest.RestName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "phone/address/restname should not be empty!")
		return
	}

	//insert rest info into database
	if err = db.InsertRest(&rest); err != nil {
		log.Printf("creat rest failed, error is %s\n", err.Error())
		if strings.Contains(err.Error(), DuplicateError) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "This phone has been registered!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Rest register failed!")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Rest %v is successfully registered!")
}
