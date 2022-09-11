package server

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/9d4/dimon/storage"
	"github.com/9d4/dimon/task"
	"github.com/gorilla/mux"
	v "github.com/spf13/viper"
)

var router *mux.Router = mux.NewRouter()

func init() {
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte{})
		}

		var t task.Task
		err = json.Unmarshal(body, &t)
		if err != nil {
			return
		}

		taskStore := task.NewStore(storage.GetDB())
		err = taskStore.Save(&t)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Unable to create new task"))
			log.Println(err)
			return
		}

		w.WriteHeader(201)
	}).Methods("POST")

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		taskStore := task.NewStore(storage.GetDB())
		tasks, err := taskStore.GetAll()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		tasksJson, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(tasksJson)
	})
}

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
	http.Serve(conn, router)
}

func Shutdown(sig chan os.Signal, conn net.Listener) {
	log.Printf("Caught signal %s: shutting down", <-sig)
	storage.Close()
	conn.Close()
	os.Exit(0)
}
