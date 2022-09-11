package task

import (
	"testing"

	"github.com/asdine/storm/v3"
)

func BenchmarkNewStore(b *testing.B) {
	db, err := storm.Open("dimon.db")
	if err != nil {
		b.Fatal(err)
	}

	// for i := 0; i < b.N; i++ {
	// 	NewStore(db)
	// }
	NewStore(db)
}
