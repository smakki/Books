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
		"Index",
		"GET",
		"/api/",
		Api_Index,
	},

	Route{
		"BookIndex",
		"GET",
		"/api/books",
		Api_BookIndex,
	},
	Route{
		"BookCreate",
		"POST",
		"/api/books",
		Api_BookCreate,
	},
	//	Route{
	//		"BookShow",
	//		"GET",
	//		"/books/{bookId}",
	//		BookShow,
	//	},
}
