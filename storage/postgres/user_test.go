package postgres

import (
	"errors"
	"fmt"
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
				Id:       "a7d8c465-8178-4455-9a8c-adb951f758c6",
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
				CreatedAt: "2021-12-20",
			},
			wantErr: false,
		},
		{
			name: "different time format testing",
			input: pb.Task{
				Id:       "e0c28933-ed82-4cbc-ac9c-437bed43200a",
				Assignee: "Abs",
				Title:    "Test",
				Summary:  "Just testing create function",
				Deadline: "01.12.2021",
				Status:   "Passed",
			},
			want: pb.Task{
				Assignee:  "Abs",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-01-12",
				Status:    "Passed",
				CreatedAt: "2021-12-20",
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

			if got.UpdatedAt != "" {
				updatedAt, err := time.Parse(time.RFC3339, got.GetUpdatedAt())
				got.UpdatedAt = updatedAt.Format("2006-01-02")
				if err != nil {
					t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
				}
			}

			got.Id = ""
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    pb.Task
		wantErr bool
	}{
		{
			name:    "not existing uuid testing",
			input:   "a7d8c465-8178-4455-9a8c-adb951f758c7",
			want:    pb.Task{},
			wantErr: true,
		},
		{
			name:  "successful",
			input: "e0c28933-ed82-4cbc-ac9c-437bed43200a",
			want: pb.Task{
				Id:        "e0c28933-ed82-4cbc-ac9c-437bed43200a",
				Assignee:  "Abs",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-01-12",
				Status:    "Passed",
				CreatedAt: "2021-12-20",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Get(tc.input)

			if tc.wantErr {
				if !reflect.DeepEqual(tc.want, got) {
					t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
				}
				if err == nil {
					t.Fatalf("%s: expected: %v, got: %v", tc.name, "no sql rows result", err)
				}
			} else {
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

				if got.UpdatedAt != "" {
					updatedAt, err := time.Parse(time.RFC3339, got.GetUpdatedAt())
					got.UpdatedAt = updatedAt.Format("2006-01-02")
					if err != nil {
						t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
					}
				}

				if !reflect.DeepEqual(tc.want, got) {
					t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
				}
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
					Id:        "a7d8c465-8178-4455-9a8c-adb951f758c6",
					Assignee:  "Lola",
					Title:     "Test",
					Summary:   "Just testing create function",
					Deadline:  "2021-12-01",
					Status:    "Passed",
					CreatedAt: "2021-12-20",
				},
				{
					Id:        "e0c28933-ed82-4cbc-ac9c-437bed43200a",
					Assignee:  "Abs",
					Title:     "Test",
					Summary:   "Just testing create function",
					Deadline:  "2021-01-12",
					Status:    "Passed",
					CreatedAt: "2021-12-20",
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotTasks, count, err := pgRepo.List(tc.page, tc.limit)
			if err != nil {
				t.Fatalf("got: %v", err)
			}
			fmt.Println(gotTasks)
			for _, task := range gotTasks {
				deadline, err := time.Parse(time.RFC3339, task.GetDeadline())
				task.Deadline = deadline.Format("2006-01-02")
				createdAt, err := time.Parse(time.RFC3339, task.GetCreatedAt())
				task.CreatedAt = createdAt.Format("2006-01-02")
				if err != nil {
					t.Fatalf("got: %v", err)
				}
			}
			if !reflect.DeepEqual(tc.want, gotTasks) && count == 2 {
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
				Id:        "a7d8c465-8178-4455-9a8c-adb951f758c6",
				Assignee:  "Lola",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-12-05",
				Status:    "Passed",
				CreatedAt: "2021-12-20",
			},
			want: pb.Task{
				Id:        "a7d8c465-8178-4455-9a8c-adb951f758c6",
				Assignee:  "Lola",
				Title:     "Test",
				Summary:   "Just testing create function",
				Deadline:  "2021-12-05",
				Status:    "Passed",
				CreatedAt: "2021-12-20",
				UpdatedAt: "2021-12-20",
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

			createdAt, err := time.Parse(time.RFC3339, got.GetCreatedAt())
			got.CreatedAt = createdAt.Format("2006-01-02")
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}

			if got.UpdatedAt != "" {
				updatedAt, err := time.Parse(time.RFC3339, got.GetUpdatedAt())
				got.UpdatedAt = updatedAt.Format("2006-01-02")
				if err != nil {
					t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
				}
			}
			if err != nil {
				t.Fatalf("got: %v", err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    error
		wantErr bool
	}{
		{
			name:    "successful",
			input:   "24465fe0-9ea1-45ce-8a7a-79c63972efe9",
			want:    nil,
			wantErr: false,
		},
		{
			name:    "delete not existing id",
			input:   "24465fe0-9ea1-45ce-8a7a-79c63972efe8",
			want:    errors.New("sql: no rows in result set"),
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := pgRepo.Delete(tc.input)
			if err == nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.want, err)
			}

			if !reflect.DeepEqual(tc.want, err) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, err)
			}
		})
	}
}

func TestTaskRepo_ListOverdue(t *testing.T) {
	tests := []struct {
		name          string
		inputDeadline string
		inputPage     int64
		inputLimit    int64
		want          []pb.Task
		wantErr       bool
	}{
		{
			name:          "successful",
			inputDeadline: "2021-12-20",
			inputPage:     1,
			inputLimit:    1,
			want: []pb.Task{
				{
					Id:        "a7d8c465-8178-4455-9a8c-adb951f758c6",
					Assignee:  "Lola",
					Title:     "Test",
					Summary:   "Just testing create function",
					Deadline:  "2021-12-05",
					Status:    "Passed",
					CreatedAt: "2021-12-20",
					UpdatedAt: "2021-12-20",
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotTasks, count, err := pgRepo.ListOverdue(tc.inputDeadline, tc.inputPage, tc.inputLimit)
			if err != nil {
				t.Fatalf("%s: expected: %v got: %v", tc.name, tc.wantErr, err)
			}

			for _, task := range gotTasks {
				deadline, err := time.Parse(time.RFC3339, task.GetDeadline())
				task.Deadline = deadline.Format("2006-01-02")
				createdAt, err := time.Parse(time.RFC3339, task.GetCreatedAt())
				task.CreatedAt = createdAt.Format("2006-01-02")
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
