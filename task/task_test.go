package task

import "testing"

func BenchmarkNewTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewTask()
	}
}
