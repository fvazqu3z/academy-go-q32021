package routes

import (
	"fmt"
	"net/http"
	"wizeline/common"

	"github.com/gorilla/mux"
)

type muxRouter struct{

}

var muxDispatcher = mux.NewRouter()

//NewMuxRouter method to create a new Mux router
func NewMuxRouter() Router {
	return &muxRouter{}
}

//GET method to manage GET requests
func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request))  {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodGet)
}

//POST method to manage POST requests
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request))  {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodPost)
}

//START to start the server
func (*muxRouter) START(port string) error {
	fmt.Printf("Server is running on port %v \n", common.Port)
	return http.ListenAndServe(port, muxDispatcher)
}


