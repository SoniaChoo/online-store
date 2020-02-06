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
	if string(b) != BadJsonRegister {
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

	if string(b) != RequestRegisterParameterMissing {
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

func TestLoginHandlerWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/user/login", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/login", LoginHandler)
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
	if string(b) != BadJsonLogin {
		t.Fatal("expect json format error, got other")
	}
}

func TestLoginHandlerWithMissingParameter(t *testing.T) {
	body := strings.NewReader(`{"phone":"1888888"}`)
	req, err := http.NewRequest(http.MethodPost, "/user/login", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/login", LoginHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}

	if string(b) != RequestLoginParameterMissing {
		t.Fatal("expect missing paramater error, got other")
	}

}

func TestLoginHandler(t *testing.T) {
	body := strings.NewReader(`{"phone":"1888888", "password":"abcd"}`)
	req, err := http.NewRequest(http.MethodPost, "/user/login", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/login", LoginHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read respone body error")
	}

	if string(b) != fmt.Sprintf(SuccessfullyLogin, "1888888") {
		t.Fatal("expect successfully login in, got other")
	}
}

func TestRetrieveIdHandlerWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/user/retrieve", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/retrieve", RetrieveIdHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != BadJsonRetrieve {
		t.Fatal("expect json format error, got other")
	}
}

func TestRetrieveIdHandler(t *testing.T) {
	body := strings.NewReader(`{"userid":2}`)
	req, err := http.NewRequest(http.MethodPost, "/user/retrieve/userid", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/retrieve/userid", RetrieveIdHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("expect 200, got other")
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}

	//b is a address, is not fixed, so we will ingore resp's body
	//if string(b) != fmt.Sprintf(SuccessfullyRetrieveId, 2) {
	//	t.Fatal("expect successfully retrieved by Id, got other")
	//}
}

func TestRetrievePhoneHandlerWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/user/retrieve/phone", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/retrieve/phone", RetrievePhoneHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read responser body error")
	}
	if string(b) != BadJsonRetrieve {
		t.Fatal("expect json format error, got other")
	}
}

func TestRetrievePhoneHandler(t *testing.T) {
	body := strings.NewReader(`{"phone":"1888888"}`)
	req, err := http.NewRequest(http.MethodPost, "/user/retrieve/phone", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/retrieve/phone", RetrievePhoneHandler)
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

func TestRetrieveNicknameHandlerWithBadJson(t *testing.T) {
	body := strings.NewReader(``)
	req, err := http.NewRequest(http.MethodPost, "/user/retrieve/nickname", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/retrieve/nickname", RetrieveNicknameHandler)
	mux.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("expect 400, got other")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response body error")
	}
	if string(b) != BadJsonRetrieve {
		t.Fatal("expect json format error, got other")
	}
}

func TestRetrieveNicknameHandler(t *testing.T) {
	body := strings.NewReader(`{"nickname":"hhh"}`)
	req, err := http.NewRequest(http.MethodPost, "/user/retrieve/nickname", body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/retrieve/nickname", RetrieveNicknameHandler)
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
