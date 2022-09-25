package task

import (
	"testing"
)

func TestEnsuresImplementStore(t *testing.T) {
	var _ Store = &store{}
}
