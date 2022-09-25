package task

import (
	"github.com/asdine/storm/v3"
)

type Store interface {
	GetAll() ([]*Task, error)
	Get(int) (*Task, error)
	Save(*Task) error
}

type store struct {
	db *storm.DB
}

func NewStore(db *storm.DB) *store {
	return &store{
		db: db,
	}
}

func (t store) GetAll() ([]*Task, error) {
	var tasks []*Task
	err := t.db.All(&tasks)

	return tasks, err
}

func (t store) Get(id int) (*Task, error) {
	var task Task
	return &task, t.db.One("ID", id, &task)
}

func (t store) Save(task *Task) error {
	return t.db.Save(task)
}
