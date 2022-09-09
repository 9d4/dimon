package task

import (
	"os/exec"
)

type Task struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Run() error {
	exc := exec.Command(t.Command, t.Args...)
	return exc.Run()
}
