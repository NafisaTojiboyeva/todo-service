package repo

import (
	pb "github.com/NafisaTojiboyeva/todo-service/genproto"
)

// TaskStorageI ...
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id string) (pb.Task, error)
	List(page, limit int64) ([]*pb.Task, int64, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id string) error
	ListOverdue(deadline string, page, limit int64) ([]*pb.Task, int64, error)
}
