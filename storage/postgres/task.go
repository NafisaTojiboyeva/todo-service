package postgres

import (
	"database/sql"
	"time"

	pb "github.com/NafisaTojiboyeva/todo-service/genproto"

	"github.com/jmoiron/sqlx"
)

type taskRepo struct {
	db *sqlx.DB
}

// NewTaskRepo ...
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task pb.Task) (pb.Task, error) {
	var id string
	err := r.db.QueryRow(`
		INSERT INTO todos(id, assignee, title, summary, deadline, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`, task.Id, task.Assignee, task.Title, task.Summary, task.Deadline, task.Status, time.Now()).Scan(&id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = r.Get(id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id string) (pb.Task, error) {
	var task pb.Task
	var updatedAt sql.NullString
	err := r.db.QueryRow(`
		SELECT id, assignee, title, summary, deadline, status, created_at, updated_at FROM todos
		WHERE id=$1 and deleted_at is null`, id).Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status, &task.CreatedAt, &updatedAt)
	if err != nil {
		return pb.Task{}, err
	}

	if !updatedAt.Valid {
		updatedAt.String = ""
	}
	task.UpdatedAt = updatedAt.String
	return task, nil
}

func (r *taskRepo) List(page, limit int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(
		`SELECT id, assignee, title, summary, deadline, status, created_at FROM todos WHERE deleted_at is null LIMIT $1 OFFSET $2`,
		limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		tasks []*pb.Task
		count int64
	)

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM todos WHERE deleted_at is null`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *taskRepo) Update(task pb.Task) (pb.Task, error) {
	result, err := r.db.Exec(`UPDATE todos SET assignee=$1, title=$2, summary=$3, deadline=$4, status=$5, updated_at=$6 WHERE id=$7 and deleted_at is null`,
		task.Assignee, task.Title, task.Summary, task.Deadline, task.Status, time.Now(), task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}

	task, err = r.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Delete(id string) error {
	result, err := r.db.Exec(`UPDATE todos SET deleted_at=$1 WHERE id=$2 and deleted_at is null`, time.Now(), id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepo) ListOverdue(deadline string, page, limit int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	time, err := time.Parse("2006-01-02", deadline)
	if err != nil {
		return nil, 0, err
	}
	rows, err := r.db.Queryx(
		`SELECT id, assignee, title, summary, deadline, status, created_at FROM todos WHERE deadline < $1 and deleted_at is null LIMIT $2 OFFSET $3 `, time, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() //nolint:errcheck

	var (
		tasks []*pb.Task
		count int64
	)

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(id) FROM todos WHERE deadline < $1 and deleted_at is null`, time).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}
