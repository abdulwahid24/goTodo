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
        	"/",
        	Index,
    	},
    	Route{
       		"TodoList",
        	"GET",
        	"/todos/",
        	TodoList,
    	},
    	Route{
        	"TodoDetail",
        	"GET",
        	"/todos/{id}/",
        	TodoDetail,
    	},
	Route{
		"TodoCreate",
		"POST",
		"/todos/",
		TodoCreate,
	},
	Route{
		"TodoUpdate",
		"PUT",
		"/todos/{id}/",
		TodoUpdate,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/todos/{id}/",
		TodoDelete,
	},
	Route{
                "TodoCreate",
                "OPTIONS",
                "/todos/",
                TodoCreate,
        },
	Route{
                "TodoUpdate",
                "OPTIONS",
                "/todos/{id}/",
                TodoUpdate,
        },
}
