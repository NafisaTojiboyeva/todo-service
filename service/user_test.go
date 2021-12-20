package service

import (
	"context"
	"fmt"
	pb "github.com/NafisaTojiboyeva/todo-service/genproto"
	"testing"
)

func TestTaskService_Create(t *testing.T) {
	_, err := client.Create(context.Background(), &pb.Task{
		Assignee:  "Lola",
		Title:     "Test",
		Summary:   "Just testing create function",
		Deadline:  "2021-12-05",
		Status:    "Passed",
		CreatedAt: "2021-12-20",
	})

	if err != nil {
		t.Error(err)
	}
}

func TestToDoService_Get(t *testing.T) {
	task, err := client.Create(context.Background(), &pb.Task{
		Assignee:  "Lola",
		Title:     "Test",
		Summary:   "Just testing create function",
		Deadline:  "2021-12-05",
		Status:    "Passed",
		CreatedAt: "2021-12-20",
	})
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Get(context.Background(), &pb.ByIdReq{
		Id: task.Id,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestToDoService_List(t *testing.T) {
	resp, err := client.List(context.Background(), &pb.ListReq{
		Limit: 2,
		Page:  1,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestToDoService_Update(t *testing.T) {
	task, err := client.Create(context.Background(), &pb.Task{
		Assignee:  "Lola",
		Title:     "Test",
		Summary:   "Just testing create function",
		Deadline:  "2021-12-05",
		Status:    "Passed",
		CreatedAt: "2021-12-20",
	})
	if err != nil {
		t.Error(err)
	}
	resp, err := client.Update(context.Background(), &pb.Task{
		Id:        task.Id,
		Assignee:  "Lola",
		Title:     "Test",
		Summary:   "Just testing create function",
		Deadline:  "2021-12-25",
		Status:    "Not send",
		CreatedAt: "2021-12-20",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestToDoService_Delete(t *testing.T) {
	task, err := client.Create(context.Background(), &pb.Task{
		Assignee:  "Lola",
		Title:     "Test",
		Summary:   "Just testing create function",
		Deadline:  "2021-12-05",
		Status:    "Passed",
		CreatedAt: "2021-12-20",
	})
	if err != nil {
		t.Error(err)
	}

	_, err = client.Delete(context.Background(), &pb.ByIdReq{
		Id: task.Id,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestToDoService_ListOverdue(t *testing.T) {
	tasks, err := client.ListOverdue(context.Background(), &pb.ByDeadlineReq{
		Deadline: "2021-12-20",
		Limit:    2,
		Page:     1,
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tasks)
}
