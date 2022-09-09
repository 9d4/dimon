package task

import (
	"bytes"
	"testing"
)

func BenchmarkNewTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewTask("Echo", "echo")
	}
}

func TestTaskRun(t *testing.T) {
	buf := bytes.NewBufferString("")

	ta := NewTask("Echoing ABC", "echo", "ABC")
	ta.cmd.Stdout = buf
	ta.Run()

	if buf.String() != "ABC\n" {
		t.Fatalf("want %s, got %s", "ABC\n", buf.String())
	}
}
