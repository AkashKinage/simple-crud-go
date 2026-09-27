package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AkashKinage/simple-crud-go/internal/model"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) (*model.Task, error)
	GetByID(ctx context.Context, id int) (*model.Task, error)
	List(ctx context.Context, filter TaskFilter) ([]*model.Task, error)
	Update(ctx context.Context, task *model.Task) (*model.Task, error)
	Delete(ctx context.Context, id int) error
}

type TaskFilter struct {
	Status string
	Priority string
	Page int
	PerPage int
}

type postgresTaskRepository struct {
    db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) TaskRepository {
    return &postgresTaskRepository{db: db}
}

func (r *postgresTaskRepository) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO tasks (title, description, status, priority)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at, updated_at`,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
	).Scan(
		&task.ID, &task.CreatedAt, &task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return task, nil
}

var ErrTaskNotFound = errors.New("task not found")

func (r *postgresTaskRepository) GetByID(ctx context.Context, id int) (*model.Task, error) {
	task := &model.Task{}

	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, title, description, status, priority, created_at, updated_at FROM tasks WHERE id = $1",
		id,
	).Scan(
		&task.ID, &task.Title, &task.Description, &task.Status,
		&task.Priority, &task.CreatedAt, &task.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTaskNotFound
		}

		return nil, err
	}

	return task, nil
}

func (r *postgresTaskRepository) List(ctx context.Context, filter TaskFilter) ([]*model.Task, error) {
	panic("not implemented")
}

func (r *postgresTaskRepository) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	panic("not implemented")
}

func (r *postgresTaskRepository) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}