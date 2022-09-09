package server

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router = mux.NewRouter()

func Serve(conn net.Listener) {
	http.Serve(conn, router)
}
