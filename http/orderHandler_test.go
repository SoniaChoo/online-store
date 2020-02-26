package http

import (
	"fmt"
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

func TestAddToCartHandlerOrderWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/order/addtocart", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/order/addtocart", AddToCartHandlerOrder)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}

	if string(b) != BadJsonAddToCart {
		t.Fatal("expect json format, got other")
	}
}

func TestAddToCartHandlerOrderWithUpdate(t *testing.T) {
	body := strings.NewReader(`{"userid":1,"dishid":1,"restid":1,"price":888,"number":9}`)
	req, err := http.NewRequest(http.MethodPost, "/order/addtocart", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/order/addtocart", AddToCartHandlerOrder)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != fmt.Sprintf(SuccessfullyAddToCart, 1) {
		t.Fatal("expect successfully add dish to cart, got other")
	}
}

func TestAddToCartHandlerOrderWithAdd(t *testing.T) {
	body := strings.NewReader(`{"userid":1,"dishid":3,"restid":1,"price":30,"number":1}`)
	req, err := http.NewRequest(http.MethodPost, "/order/addtocart", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/order/addtocart", AddToCartHandlerOrder)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != fmt.Sprintf(SuccessfullyAddToCart, 3) {
		t.Fatal("expect successfully add dish to cart, got other")
	}
}

func TestDeleteDishInCartHandlerOrderWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/order/deletedishincart", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/order/deletedishincart", DeleteDishInCartHandlerOrder)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != BadJsonDeleteDishInCart {
		t.Fatal("expect json format, got other")
	}
}

func TestDeleteDishInCartHandlerOrderWithDelete(t *testing.T) {
	body := strings.NewReader(`{"detailid":1,"dishid":1,"restid":1,"orderid":1,"price":888,"number":0,"status":-1}`)
	req, err := http.NewRequest(http.MethodPost, "/order/deletedishincart", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/order/deletedishincart", DeleteDishInCartHandlerOrder)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}

	if string(b) != fmt.Sprintf(SuccessfullyDeleteOrUpdateInCart, 1) {
		t.Fatal("expect successfully delete dish in cart, got other")
	}
}
