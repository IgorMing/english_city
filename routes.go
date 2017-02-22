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
		indexHandler,
	},
	Route{
		"Rooms",
		"GET",
		"/rooms",
		roomsHandler,
	},
	Route{
		"RoomById",
		"GET",
		"/rooms/{id}",
		roomsByIDHandler,
	},
	Route{
		"AddRoom",
		"POST",
		"/rooms",
		addRoomHandler,
	},
	Route{
		"DeleteRoom",
		"DELETE",
		"/rooms/{id}",
		deleteRoomHandler,
	},
}
