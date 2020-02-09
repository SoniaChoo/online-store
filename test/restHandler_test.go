package test

import (
	"fmt"
	http2 "github.com/SoniaChoo/online-store/http"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandlerRestWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/rest/register", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/rest/register", http2.RegisterHandlerRest)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != http2.BadJsonRest {
		t.Fatal("expect json format error, got other")
	}
}

func TestRegisterHandlerRestWithMissParameter(t *testing.T) {
	body := strings.NewReader(`{"phone":"1999999", "restname":"yuyuyu"}`)
	req, err := http.NewRequest(http.MethodPost, "/rest/register", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/rest/register", http2.RegisterHandlerRest)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != http2.RequestParameterMissingRest {
		t.Fatal("expect missing parameter error, got other")
	}
}

func TestRegisterHandlerRest(t *testing.T) {
	body := strings.NewReader(`{"phone":"1999999", "restname":"yuyuyu", "address":"yangxin", "userid":2}`)
	req, err := http.NewRequest(http.MethodPost, "/rest/register", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/rest/register", http2.RegisterHandlerRest)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != fmt.Sprintf(http2.SuccessfullyRegisterRest, "1999999") {
		t.Fatal("expect successfully register, got other")
	}
}

func TestShowDishesHandlerRestWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/rest/showdishes", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/rest/showdishes", http2.ShowDishesHandlerRest)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != http2.BadJsonRest {
		t.Fatal("expect json format error, got other")
	}
}

func TestShowDishesHandlerRest(t *testing.T) {
	body := strings.NewReader(`{"restid":1}`)
	req, err := http.NewRequest(http.MethodPost, "/rest/showdishes", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/rest/showdishes", http2.ShowDishesHandlerRest)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
}

func TestRetrieveHandlerRestWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/rest/retrieve", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/rest/retrieve", http2.RegisterHandlerRest)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != http2.BadJsonRest {
		t.Fatal("expect json format error, got other")
	}
}

func TestRetrieveHandlerRest(t *testing.T) {
	body := strings.NewReader(`{"restname":"yuyuyu"}`)
	req, err := http.NewRequest(http.MethodPost, "/rest/retrieve", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/rest/retrieve", http2.RetrieveHandlerRest)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
}
