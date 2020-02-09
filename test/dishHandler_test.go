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

func TestDetailHandlerDishWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/dish/detail", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/detail", http2.DetailHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != http2.BadJsonDishDetail {
		t.Fatal("expect json format error, got other")
	}
}

func TestDetailHandlerDish(t *testing.T) {
	body := strings.NewReader(`{"dishid":1}`)
	req, err := http.NewRequest(http.MethodPost, "/dish/detail", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/detail", http2.DetailHandlerDish)
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

func TestAddHandlerDishWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/dish/add", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/add", http2.AddHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != http2.BadJsonDishAdd {
		t.Fatal("expect json format error, got other")
	}
}

func TestAddHandlerDishWithMissParameter(t *testing.T) {
	body := strings.NewReader(`{"restid":1,"price":30,"dishname":"maocai","description":"sichuan cai, is delicious"}`)
	req, err := http.NewRequest(http.MethodPost, "/dish/add", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/add", http2.AddHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("expect missing parameter error, got other")
	}
	if string(b) != http2.RequestParameterMissingDish {
		t.Fatal("expect missing parameter error, got other")
	}
}

func TestAddHandlerDish(t *testing.T) {
	body := strings.NewReader(`{"restid":1,"price":30,"dishname":"maocai","description":"sichuan cai, is delicious", "stock":10}`)
	req, err := http.NewRequest(http.MethodPost, "/dish/add", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/add", http2.AddHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != fmt.Sprintf(http2.SuccessfullyAddDish, "maocai") {
		t.Fatal("expect successfully add, got other")
	}
}

func TestUpdateHandlerDishWithBadJSon(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/dish/update", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/update", http2.UpdateHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != http2.BadJsonDishUpdate {
		t.Fatal("expect json format error, got other")
	}
}

func TestUpdateHandlerDishWithNotAvailable(t *testing.T) {
	body := strings.NewReader(`{"price":0,"dishname":"maocai","description":"sichuan cai, is delicious","stock":10,"dishid":1}`)
	req, err := http.NewRequest(http.MethodPost, "/dish/update", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/update", http2.UpdateHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatal("expect 500, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != http2.RequestUpdateNotAvailable {
		t.Fatal("expect parameter not available error, got other")
	}
}

func TestUpdateHandlerDish(t *testing.T) {
	body := strings.NewReader(`{"price":888,"dishname":"maocai","description":"sichuan cai, is delicious","stock":10,"dishid":1}`)
	req, err := http.NewRequest(http.MethodPost, "/dish/update", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/update", http2.UpdateHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != fmt.Sprintf(http2.SuccessfullyUpdateDish, 1) {
		t.Fatal("expect successfully update dish, got other")
	}
}

func TestSearchByDishNameHandlerDishWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/dish/search/name", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/search/name", http2.SearchByDishNameHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}

	if string(b) != http2.BadJsonDishSearch {
		t.Fatal("expect json format error, got other")
	}
}

func TestSearchByDishNameHandlerDish(t *testing.T) {
	body := strings.NewReader(`{"dishname":"mao"}`)
	req, err := http.NewRequest(http.MethodPost, "/dish/search/name", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/dish/search/name", http2.SearchByDishNameHandlerDish)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	if _, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fatal("read response body error")
	}
}
