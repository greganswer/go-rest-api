package main

import "github.com/julienschmidt/httprouter"

// Route is composite of route info
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

// Routes is a slice of Route
type Routes []Route

// AllRoutes returns the list of routes
func AllRoutes() Routes {
	return Routes{
		Route{"Index", "GET", "/", Index},
		Route{"BookIndex", "GET", "/books", BookIndex},
		Route{"BookShow", "GET", "/books/:isdn", BookShow},
		Route{"BookShow", "POST", "/books", BookCreate},
	}
}
