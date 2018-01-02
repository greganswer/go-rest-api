package main

// Book is a record
type Book struct {
	ISDN   string `json:"isdn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

// A map to store the books with the ISDN as the key
// This acts as the storage in lieu of an actual database
var bookstore = make(map[string]*Book)
