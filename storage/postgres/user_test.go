package postgres

import (
	pb "github.com/NafisaTojiboyeva/todo-service/genproto"
	"reflect"
	"testing"
	"time"
)

func TestTaskRepo_Create(t *testing.T) {
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
				Assignee: "Lola",
				Title:    "Test",
				Summary:  "Just testing create function",
				Deadline: "2021-12-01",
				Status:   "Passed",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Create(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			got.Id = 0
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    pb.Task
		wantErr bool
	}{
		{
			name:  "successful",
			input: 1,
			want: pb.Task{
				Id:       1,
				Assignee: "jack",
				Title:    "homework",
				Summary:  "solve the arithmetic problem",
				Deadline: "2021-12-15",
				Status:   "not send",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Get(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			deadline, err := time.Parse(time.RFC3339, got.GetDeadline())
			got.Deadline = deadline.Format("2006-01-02")
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_List(t *testing.T) {
	tests := []struct {
		name    string
		page    int64
		limit   int64
		want    []pb.Task
		wantErr bool
	}{
		{
			name:  "successful",
			page:  1,
			limit: 2,
			want: []pb.Task{
				{
					Id:       1,
					Assignee: "jack",
					Title:    "homework",
					Summary:  "solve the arithmetic problem",
					Deadline: "2021-12-15",
					Status:   "not send",
				},
				{
					Id:       2,
					Assignee: "john",
					Title:    "classwork",
					Summary:  "write recursive function",
					Deadline: "2021-12-20",
					Status:   "passed",
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotTasks, count, err := pgRepo.List(tc.page, tc.limit)
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}

			for _, task := range gotTasks {
				deadline, err := time.Parse(time.RFC3339, task.GetDeadline())
				task.Deadline = deadline.Format("2006-01-02")
				if err != nil {
					t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
				}
			}
			if !reflect.DeepEqual(tc.want, gotTasks) && count == 4 {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, gotTasks)
			}
		})
	}
}

func TestTaskRepo_Update(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Task
		want    pb.Task
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Task{
				Id:       1,
				Assignee: "JACk",
				Title:    "Something",
				Summary:  "There isn't any summary",
				Deadline: "2021-12-18",
				Status:   "Checking",
			},
			want: pb.Task{
				Id:       1,
				Assignee: "JACk",
				Title:    "Something",
				Summary:  "There isn't any summary",
				Deadline: "2021-12-18",
				Status:   "Checking",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Update(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			deadline, err := time.Parse(time.RFC3339, got.GetDeadline())
			got.Deadline = deadline.Format("2006-01-02")
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	tests := []struct{
		name string
		input int64
		want
	}
}
