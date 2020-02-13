package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowCartHandlerOrderWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/order/showcart", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/order/showcart", ShowCartHandlerOrder)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != BadJsonShowCart {
		t.Fatal("expect json format error, got other")
	}
}

func TestShowCartHandlerOrder(t *testing.T) {
	body := strings.NewReader(`{"userid":1}`)
	req, err := http.NewRequest(http.MethodPost, "/order/showcart", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/order/showcart", ShowCartHandlerOrder)
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
