package storage

import (
	"github.com/NafisaTojiboyeva/todo-service/storage/postgres"
	"github.com/NafisaTojiboyeva/todo-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Task() repo.TaskStorageI
}

type storagePg struct {
	db       *sqlx.DB
	taskRepo repo.TaskStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		taskRepo: postgres.NewTaskRepo(db),
	}
}

func (s storagePg) Task() repo.TaskStorageI {
	return s.taskRepo
}
