package task

import (
	"os/exec"
)

type Task struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Command string   `json:"command"`
	Args    []string `json:"args"`

	cmd *exec.Cmd
}

func NewTask(name, command string, args ...string) *Task {
	return &Task{
		Name:    name,
		Command: command,
		Args:    args,
		cmd:     parseCommandArgs(command, args),
	}
}

func parseCommandArgs(command string, args []string) *exec.Cmd {
	return exec.Command(command, args...)
}

func (t *Task) Run() error {
	return t.cmd.Run()
}
