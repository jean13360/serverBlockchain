package main

import (
	"blockchain/services"
	"net/http"
)

//Route model pour créer une route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes liste de route
type Routes []Route

var routes = Routes{
	Route{
		"Status",
		"GET",
		"/",
		Services.Status,
	}, Route{
		"Block",
		"POST",
		"/block",
		Services.CreateBlock,
	},
}
