package process

import (
	"bytes"
	"testing"
)

func TestNewProcess(t *testing.T) {
	NewProcess(0, "ls", "-lah")
}

func TestProcessRun(t *testing.T) {
	buf := bytes.NewBufferString("")
	p := NewProcess(0, "echo", "ABC")
	p.Cmd.Stdout = buf
	p.Run()

	if buf.String() != "ABC\n" {
		t.Fatalf("want %s, got %s", "ABC\n", buf.String())
	}

	if p.IsRunning() {
		t.Fatal("process p should has been stopped")
	}
}

func TestProcessKill(t *testing.T) {
	buf := bytes.NewBufferString("")

	p := NewProcess(0, "sleep", "10")
	p.Cmd.Stdout = buf

	err := p.Start()
	if err != nil {
		t.Fatal(err)
	}

	if p.Cmd.ProcessState != nil {
		t.Fatal("process p should be running")
	}

	p.Kill()

	if p.Cmd.ProcessState != nil && p.Cmd.ProcessState.Exited() {
		t.Fatal("process p should has been stopped")
	}
}
