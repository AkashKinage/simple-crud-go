package service

import (
	"context"
	"errors"

	"github.com/AkashKinage/simple-crud-go/internal/model"
	"github.com/AkashKinage/simple-crud-go/internal/repository"
)

type TaskService interface {
	Create(ctx context.Context, task *model.Task) (*model.Task, error)
	GetByID(ctx context.Context, id int) (*model.Task, error)
	List(ctx context.Context, filter repository.TaskFilter) ([]*model.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

var ErrInvalidTitle = errors.New("title is required")
var ErrInvalidStatus = errors.New("invalid status")
var ErrInvalidPriority = errors.New("invalid priority")

var ErrInvalidPerPage = errors.New("per_page must not exceed 100")

func (s *taskService) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	if task.Title == "" {
		return nil, ErrInvalidTitle
	}

	if task.Status == "" {
		task.Status = "pending"
	} else if !model.IsValidStatus(task.Status) {
		return nil, ErrInvalidStatus
	}

	if task.Priority == "" {
		task.Priority = "low"
	} else if !model.IsValidPriority(task.Priority) {
		return nil, ErrInvalidPriority
	}

	return s.repo.Create(ctx, task)
}

func (s *taskService) GetByID(ctx context.Context, id int) (*model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *taskService) List(ctx context.Context, filter repository.TaskFilter) ([]*model.Task, error) {
	if filter.Status != "" && !model.IsValidStatus(filter.Status) {
		return nil, ErrInvalidStatus
	}

	if filter.Priority != "" && !model.IsValidPriority(filter.Priority) {
		return nil, ErrInvalidPriority
	}

	if filter.PerPage != 0 && filter.PerPage > 100 {
		return nil, ErrInvalidPerPage
	}

	return s.repo.List(ctx, filter)
}
