package server

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	v "github.com/spf13/viper"
)

var router *mux.Router = mux.NewRouter()

func quickSetup() {
	err := os.MkdirAll(v.GetString("socketdir"), os.ModePerm)
	checkErr(err)

	os.Remove(v.GetString("socketpath"))
}

func listenSocket() {
	listener, err := net.Listen("unix", v.GetString("socketpath"))
	checkErr(err)

	err = os.Chmod(v.GetString("socketpath"), os.FileMode(v.GetInt32("socketmask")))
	checkErr(err)

	log.Println("Listening in", listener.Addr().String())
	Serve(listener)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func Start() {
	quickSetup()
	listenSocket()
}

func Serve(conn net.Listener) {
	http.Serve(conn, router)
}
