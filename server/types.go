package server

import (
	"github.com/9d4/dimon/process"
	"github.com/9d4/dimon/task"
)

type Process struct {
	process.Process

	// Task name
	Task task.Task

	// The running command
	Run string

	// Determines running or not
	Status bool
}
