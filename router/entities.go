package router

import "net/http"

type routeStruct struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routesStruct []routeStruct

var env string
var routes routesStruct

var Value int

type Resp struct {
	Message string `json:"message"`
}

type Req struct {
	Value int `json:"value"`
}
