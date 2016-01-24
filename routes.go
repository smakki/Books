package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ApiIndex",
		"GET",
		"/api/",
		Api_Index,
	},

	Route{
		"ApiBookIndex",
		"GET",
		"/api/books",
		Api_BookIndex,
	},
	Route{
		"ApiBookCreate",
		"POST",
		"/api/books",
		Api_BookCreate,
	},
	Route{
		"ApiBookShow",
		"GET",
		"/api/books/{bookId}",
		Api_BookShow,
	},
}
