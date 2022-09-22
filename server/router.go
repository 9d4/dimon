package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/9d4/dimon/process"
	"github.com/9d4/dimon/storage"
	"github.com/9d4/dimon/task"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

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

		// validate
		if t.Name == "" {
			w.WriteHeader(400)
			return
		}

		if t.Command == "" {
			w.WriteHeader(400)
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
		w.Write(body)
	}).Methods("POST")

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		taskStore := task.NewStore(storage.GetDB())
		tasks, err := taskStore.GetAll()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		var custTasks []Task
		for _, t := range tasks {
			var ta Task
			ta.Task = *t
			ta.CommandArgs = fmt.Sprintf("%s %s", t.Command, strings.Join(t.Args, ""))
			custTasks = append(custTasks, ta)
		}

		tasksJson, err := json.Marshal(custTasks)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(tasksJson)
	})

	router.HandleFunc("/tasks/{taskid}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		taskID, err := strconv.Atoi(vars["taskid"])
		if err != nil {
			w.WriteHeader(400)
			return
		}

		task, err := (task.NewStore(storage.GetDB())).Get(taskID)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		// create new process based on selected task
		proc := process.NewProcess(taskID, task.Command, task.Args...)
		process.SaveProcess(proc)

		proc.Start()

		buf, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(500)
		}

		w.WriteHeader(200)
		w.Write(buf)
	}).Methods("POST")

	router.HandleFunc("/processes", func(w http.ResponseWriter, r *http.Request) {
		procs := process.GetAll()

		var custProcs []Process
		taskStore := task.NewStore(storage.GetDB())

		for _, p := range procs {
			var cp Process

			t, err := taskStore.Get(p.TaskID)
			if err != nil {
				continue
			}

			cp.Task = *t
			cp.Run = fmt.Sprintf("%s %s", t.Command, strings.Join(t.Args, " "))
			cp.Status = p.IsRunning()
			cp.PID = p.Cmd.Process.Pid

			custProcs = append(custProcs, cp)
		}

		buf, err := json.Marshal(custProcs)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(buf)
	}).Methods("GET")

	return router
}
