package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandlerWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/user/register", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/register", RegisterHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	fmt.Println(string(b))
	if string(b) != BadJson {
		t.Fatal("expect json format error, got other")
	}
}

func TestRegisterHandlerWithMissingParameter(t *testing.T) {
	body := strings.NewReader(`{"phone":"1888888", "password":"zzz"}`)
	req, err := http.NewRequest(http.MethodPost, "/user/register", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/register", RegisterHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}

	if string(b) != RequestParameterMissing {
		t.Fatal("expect missing paramater error, got other")
	}
}

func TestRegisterHandler(t *testing.T) {
	body := strings.NewReader(`{"phone":"1888888", "nickname":"zz", "password":"abcd"}`)
	req, err := http.NewRequest(http.MethodPost, "/user/register", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/register", RegisterHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}

	if string(b) != fmt.Sprintf(SuccessfullyRegister, "1888888") {
		t.Fatal("expect successfully registered, got other")
	}
}
