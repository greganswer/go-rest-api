package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestBookIndex(t *testing.T) {
	testBook := &Book{
		ISDN:   "111",
		Title:  "test title",
		Author: "test author",
		Pages:  42,
	}
	bookstore["111"] = testBook
	req1, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr1 := newRequestRecorder(req1, "GET", "/books", BookIndex)
	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}
	exp := "{\"meta\":null,\"data\":[{\"isdn\":\"111\",\"title\":\"test title\",\"author\":\"test author\",\"pages\":42}]}\n"
	actual := rr1.Body.String()
	if actual != exp {
		t.Fatalf("\nResponse body does not match: \nExpected: '%s' \n  Actual: '%s'", exp, actual)
	}
}

// Mocks a handler and returns a httptest.ResponseRecorder
func newRequestRecorder(req *http.Request, method string, path string, handler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, path, handler)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
