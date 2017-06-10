package main

import "net/http"

// Route is the structure for reach route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes stores all routes in a slice
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1",
		Index,
	},
	Route{
		"Fish",
		"POST",
		"/v1/fish",
		Fishy,
	},
	Route{
		"Websocket",
		"GET",
		"/v1/ws",
		OpenWS,
	},
	Route{
		"GetLocation",
		"GET",
		"/v1/location/{userID}",
		Location,
	},
	Route{
		"SetLocation",
		"PUT",
		"/v1/location/{userID}/{loc}",
		Location,
	},
	Route{
		"GetInventory",
		"GET",
		"/v1/inventory/{userID}",
		Inventory,
	},
	Route{
		"SetItem",
		"POST",
		"/v1/inventory/{userID}",
		BuyItem,
	},
	Route{
		"Blacklist",
		"GET",
		"/v1/blacklist/{userID}",
		Blacklist,
	},
	Route{
		"Unblacklist",
		"DELETE",
		"/v1/blacklist/{userID}",
		Unblacklist,
	},
}
