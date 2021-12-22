package service

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/NafisaTojiboyeva/todo-service/genproto"
	"reflect"
	"testing"
	"time"
)

func TestTaskService_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Task
		want    pb.Task
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Task{
				Assignee: "Lola",
				Title:    "Test",
				Summary:  "Just testing create function",
				Deadline: "2021-12-01",
				Status:   "Passed",
			},
			want: pb.Task{
				Assignee:  "Lola",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-12-01",
				Status:    "Passed",
				CreatedAt: "2021-12-22",
			},
			wantErr: false,
		},
		{
			name: "different time format testing",
			input: pb.Task{
				Assignee: "Abs",
				Title:    "Test",
				Summary:  "Just testing create function",
				Deadline: "12.25.2021",
				Status:   "Passed",
			},
			want: pb.Task{
				Assignee:  "Abs",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-12-25",
				Status:    "Passed",
				CreatedAt: "2021-12-22",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Create(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}

			// formatting times
			deadline, err := time.Parse(time.RFC3339, got.GetDeadline())
			got.Deadline = deadline.Format("2006-01-02")
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			createdAt, err := time.Parse(time.RFC3339, got.GetCreatedAt())
			got.CreatedAt = createdAt.Format("2006-01-02")

			got.Id = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestToDoService_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.ByIdReq
		want    pb.Task
		wantErr bool
	}{
		{
			name:  "successful",
			input: pb.ByIdReq{Id: "24465fe0-9ea1-45ce-8a7a-79c63972efe9"},
			want: pb.Task{
				Assignee:  "Lola",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-12-01",
				Status:    "Passed",
				CreatedAt: "2021-12-21",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Get(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}

			// formatting times
			deadline, err := time.Parse(time.RFC3339, got.GetDeadline())
			got.Deadline = deadline.Format("2006-01-02")
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			createdAt, err := time.Parse(time.RFC3339, got.GetCreatedAt())
			got.CreatedAt = createdAt.Format("2006-01-02")

			got.Id = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestToDoService_List(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.ListReq
		want    pb.ListResp
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.ListReq{
				Limit: 1,
				Page:  1,
			},
			want: pb.ListResp{
				Tasks: []*pb.Task{
					{
						Id:        "24465fe0-9ea1-45ce-8a7a-79c63972efe9",
						Assignee:  "Lola",
						Title:     "Test",
						Summary:   "Just testing create function",
						Deadline:  "2021-12-01",
						Status:    "Passed",
						CreatedAt: "2021-12-21",
					},
				},
				Count: 2,
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.List(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}

			// formatting times
			for _, task := range got.Tasks {
				deadline, err := time.Parse(time.RFC3339, task.Deadline)
				task.Deadline = deadline.Format("2006-01-02")
				if err != nil {
					t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
				}
				createdAt, err := time.Parse(time.RFC3339, task.GetCreatedAt())
				task.CreatedAt = createdAt.Format("2006-01-02")
			}

			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestToDoService_Update(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Task
		want    pb.Task
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Task{
				Id:        "24465fe0-9ea1-45ce-8a7a-79c63972efe9",
				Assignee:  "Lola",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-12-05",
				Status:    "Passed",
				CreatedAt: "2021-12-21",
			},
			want: pb.Task{
				Id:        "24465fe0-9ea1-45ce-8a7a-79c63972efe9",
				Assignee:  "Lola",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-12-05",
				Status:    "Passed",
				CreatedAt: "2021-12-21",
				UpdatedAt: "2021-12-21",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Update(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}

			// formatting times
			deadline, err := time.Parse(time.RFC3339, got.GetDeadline())
			got.Deadline = deadline.Format("2006-01-02")
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			createdAt, err := time.Parse(time.RFC3339, got.GetCreatedAt())
			got.CreatedAt = createdAt.Format("2006-01-02")
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			updatedAt, err := time.Parse(time.RFC3339, got.GetUpdatedAt())
			got.UpdatedAt = updatedAt.Format("2006-01-02")
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}

			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestToDoService_Delete(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.ByIdReq
		want    error
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.ByIdReq{
				Id: "def039c9-e169-4301-86d7-36d346d5502e",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "delete not existing id",
			input: pb.ByIdReq{
				Id: "24465fe0-9ea1-45ce-8a7a-79c63972efe8",
			},
			want:    errors.New("rpc error: code = Internal desc = failed to delete task"),
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.Delete(context.Background(), &tc.input)
			fmt.Println(err)
			if err != nil {
				if !reflect.DeepEqual(tc.want, err) {
					t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, err)
				}
			}

		})
	}
}

func TestToDoService_ListOverdue(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.ByDeadlineReq
		want    pb.ListResp
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.ByDeadlineReq{
				Deadline: "2021-12-06",
				Limit:    2,
				Page:     1,
			},
			want: pb.ListResp{
				Tasks: []*pb.Task{
					{
						Id:        "2128d9a8-bc96-4fcf-85e4-9a6e4493b1c2",
						Assignee:  "Lola",
						Title:     "Test",
						Summary:   "Just testing create function",
						Deadline:  "2021-12-01",
						Status:    "Passed",
						CreatedAt: "2021-12-22",
					},
				},
				Count: 1,
			},
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.ListOverdue(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}

			// formatting times
			for _, task := range got.Tasks {
				deadline, err := time.Parse(time.RFC3339, task.Deadline)
				task.Deadline = deadline.Format("2006-01-02")
				if err != nil {
					t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
				}
				createdAt, err := time.Parse(time.RFC3339, task.GetCreatedAt())
				task.CreatedAt = createdAt.Format("2006-01-02")
			}

			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}
