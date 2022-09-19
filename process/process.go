package process

import (
	"os"
	"os/exec"
)

// Process is wrapper for os/exec. Process used to run process as background process.
type Process struct {
	TaskID int       `json:"taskID"`
	Cmd    *exec.Cmd `json:"-"`
}

func NewProcess(taskid int, command string, args ...string) *Process {
	p := &Process{
		TaskID: taskid,
		Cmd:    parseCommandArgs(command, args),
	}
	p.Cmd.Stdout = os.Stdout
	return p
}

func (p *Process) Start() error {
	return p.Cmd.Start()
}

func (p *Process) Run() error {
	return p.Cmd.Run()
}

func (p *Process) Kill() error {
	return p.Cmd.Process.Kill()
}

func (p *Process) IsRunning() bool {
	if p.Cmd.ProcessState != nil {
		return !p.Cmd.ProcessState.Exited()
	}

	return true
}

func parseCommandArgs(command string, args []string) *exec.Cmd {
	return exec.Command(command, args...)
}
