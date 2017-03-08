package routes

import (
	"net/http"

	"github.com/IgorMing/englishcity/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var roomRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.IndexHandler,
	},
	Route{
		"Rooms",
		"GET",
		"/rooms",
		handlers.RoomsHandler,
	},
	Route{
		"RoomById",
		"GET",
		"/rooms/{id}",
		handlers.RoomsByIDHandler,
	},
	Route{
		"AddRoom",
		"POST",
		"/rooms",
		handlers.AddRoomHandler,
	},
	Route{
		"UpdateRoom",
		"PUT",
		"/rooms/{id}",
		handlers.UpdateRoomHandler,
	},
	Route{
		"DeleteRoom",
		"DELETE",
		"/rooms/{id}",
		handlers.DeleteRoomHandler,
	},
}
