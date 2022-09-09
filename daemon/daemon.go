package daemon

import (
	"log"
	"net"
	"os"

	"github.com/9d4/dimon/server"
	v "github.com/spf13/viper"
)

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
	server.Serve(listener)
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
