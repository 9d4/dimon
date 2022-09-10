package process

import (
	"os"
	"os/exec"
)

// Process is wrapper for os/exec. Process used to run process as background process.
type Process struct {
	Cmd *exec.Cmd
}

func NewProcess(command string, args ...string) *Process {
	p := &Process{
		Cmd: parseCommandArgs(command, args),
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
	return !p.Cmd.ProcessState.Exited()
}

func parseCommandArgs(command string, args []string) *exec.Cmd {
	return exec.Command(command, args...)
}
