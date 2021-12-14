package service

import (
	pb "app/genproto"
	l "app/pkg/logger"
	"app/storage"
	"context"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ToDoService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewToDoService
func NewToDoService(db *sqlx.DB, log l.Logger) *ToDoService {
	return &ToDoService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *ToDoService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	task, err := s.storage.Task().Create(*req)
	if err != nil {
		s.logger.Error("failed to create task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &task, nil
}

func (s *ToDoService) Get(ctx context.Context, req *pb.ByIdReq) (*pb.Task, error) {
	task, err := s.storage.Task().Get(req.GetId())
	if err != nil {
		s.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &task, nil
}

func (s *ToDoService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	task, count, err := s.storage.Task().List(req.Page, req.Limit)
	if err != nil {
		s.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &pb.ListResp{
		Tasks: task,
		Count: count,
	}, nil
}

func (s *ToDoService) Update(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	task, err := s.storage.Task().Update(*req)
	if err != nil {
		s.logger.Error("failed to update task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &task, nil
}

func (s *ToDoService) Delete(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := s.storage.Task().Delete(req.Id)
	if err != nil {
		s.logger.Error("failed to delete task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &pb.EmptyResp{}, nil
}

func (s *ToDoService) ListOverdue(ctx context.Context, req *pb.ByDeadlineReq) (*pb.ListResp, error) {
	tasks, count, err := s.storage.Task().ListOverdue(req.Deadline)
	if err != nil {
		s.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &pb.ListResp{
		Tasks: tasks,
		Count: count,
	}, nil
}
