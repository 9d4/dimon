package task

import "os/exec"

type Task struct {
	Name    string
	Command string
	Args    []string
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Run() error {
	exc := exec.Command(t.Command, t.Args...)
	return exc.Run()
}
