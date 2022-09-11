package task

import (
	"github.com/asdine/storm/v3"
)

type Store struct {
	db *storm.DB
}

func NewStore(db *storm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (t Store) GetAll() ([]*Task, error) {
	var tasks []*Task
	err := t.db.All(&tasks)

	return tasks, err
}

func (t Store) Save(task *Task) error {
	return t.db.Save(task)
}
