package daemon

import (
	"bufio"
	"log"
	"net"
	"os"

	v "github.com/spf13/viper"
)

func quickSetup() {
	err := os.MkdirAll(v.GetString("socketdir"), os.ModePerm)
	checkErr(err)

	err = os.Remove(v.GetString("socketpath"))
	checkErr(err)
}

func listenSocket() {
	listener, err := net.Listen("unix", v.GetString("socketpath"))
	checkErr(err)

	err = os.Chmod(v.GetString("socketpath"), os.FileMode(v.GetInt32("socketmask")))
	checkErr(err)

	log.Println("Listening in", listener.Addr().String())

	for {
		go handleConn(listener.Accept())
	}
}

func handleConn(conn net.Conn, err error) {
	bufreader := bufio.NewReader(conn)

	for {
		conn.Write([]byte(">>> "))

		_, err := bufreader.ReadBytes('\n')
		if err != nil {
			conn.Close()
			break
		}
	}
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
