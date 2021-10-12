package entities

import(
		"net/http"
)

type Route struct{
	Name string							`json:"name"`
	Method string						`json:"method"`
	Path string							`json:"pattern"`
	HandleFunc http.HandlerFunc			`json:"handleFunc"`
}

type Routes [] Route