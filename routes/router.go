package routes

import (
	"net/http"
)

//Router interface to manage server operations
type Router interface{
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	START(port string) error
}



