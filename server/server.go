package server

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/9d4/dimon/storage"
	v "github.com/spf13/viper"
)

func quickSetup() {
	if _, err := os.Stat(v.GetString("socketpath")); os.IsNotExist(err) {
		err := os.MkdirAll(v.GetString("socketpath"), os.ModePerm)
		checkErr(err)
	}

	os.Remove(v.GetString("socketpath"))
}

func listenSocket() {
	listener, err := net.Listen("unix", v.GetString("socketpath"))
	checkErr(err)

	err = os.Chmod(v.GetString("socketpath"), os.FileMode(v.GetInt32("socketmask")))
	checkErr(err)

	var sigc = make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go Shutdown(sigc, listener)

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
	http.Serve(conn, NewRouter())
}

func Shutdown(sig chan os.Signal, conn net.Listener) {
	log.Printf("Caught signal %s: shutting down", <-sig)
	storage.Close()
	conn.Close()
	os.Exit(0)
}
