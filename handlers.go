package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index is the home route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func BookCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	book := &Book{}
	if err := populateModelFromHandler(w, r, params, book); err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	bookstore[book.ISDN] = book
	writeOKResponse(w, book)
}

// BookIndex is the handler for the books index action
// GET /books
func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := []*Book{}
	for _, book := range bookstore {
		books = append(books, book)
	}
	writeOKResponse(w, books)
}

// BookShow is the handler for the books Show action
// GET /books/:isdn
func BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isdn := params.ByName("isdn")
	book, ok := bookstore[isdn]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !ok {
		writeErrorResponse(w, http.StatusNotFound, "Record Not Found")
		return
	}
	writeOKResponse(w, book)
}

// writeOKResponse sends a JSON respose with status OK (200)
func writeOKResponse(w http.ResponseWriter, d interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JSONResponse{Data: d}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// writeErrorResponse sends a JSON error response with the given status code
func writeErrorResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	errResp := &JSONErrorResponse{Error: &APIError{Status: int16(code), Title: msg}}
	json.NewEncoder(w).Encode(errResp)
}

// populateModelFromHandler populates a model from the params in the Handler
func populateModelFromHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	return json.Unmarshal(body, model)
}
