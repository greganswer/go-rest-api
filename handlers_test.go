package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestBookShow(t *testing.T) {
	t.Log("When the books' ISDN does not exist")
	req1, err := http.NewRequest("GET", "/books/1234", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr1 := newRequestRecorder(req1, "GET", "/books/:isdn", BookShow)
	expCode := 404
	if rr1.Code != expCode {
		t.Fatalf("Expected response code to be %d, got %d", expCode, rr1.Code)
	}
	expRes1 := "{\"error\":{\"status\":404,\"title\":\"Record Not Found\"}}\n"
	actual1 := rr1.Body.String()
	if actual1 != expRes1 {
		t.Fatalf("\nResponse body does not match: \nExpected: '%s' \n  Actual: '%s'", expRes1, actual1)
	}
	t.Log("When the book exists")
	// Create an entry of the book to the bookstore map
	testBook := &Book{
		ISDN:   "111",
		Title:  "test title",
		Author: "test author",
		Pages:  42,
	}
	bookstore["111"] = testBook
	req2, err := http.NewRequest("GET", "/books/111", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr2 := newRequestRecorder(req2, "GET", "/books/:isdn", BookShow)
	expCode2 := 200
	if rr2.Code != expCode2 {
		t.Fatalf("Expected response code to be %d, got %d", expCode, rr2.Code)
	}
	// expected response
	expRes2 := "{\"meta\":null,\"data\":{\"isdn\":\"111\",\"title\":\"test title\",\"author\":\"test author\",\"pages\":42}}\n"
	actual2 := rr2.Body.String()
	if actual2 != expRes2 {
		t.Fatalf("\nResponse body does not match: \nExpected: '%s' \n  Actual: '%s'", expRes2, actual2)
	}
}
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
