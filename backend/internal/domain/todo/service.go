package todo

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrTodoNotFound     = errors.New("todo not found")
	ErrInvalidTitle     = errors.New("title cannot be empty and must be less than 200 characters")
	ErrInvalidPriority  = errors.New("priority must be low, medium, or high")
	ErrDescriptionTooLong = errors.New("description cannot exceed 1000 characters")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTodo(ctx context.Context, title string, description *string, priority Priority) (*Todo, error) {
	if err := s.validateTitle(title); err != nil {
		return nil, err
	}

	if err := s.validateDescription(description); err != nil {
		return nil, err
	}

	if !priority.IsValid() {
		return nil, ErrInvalidPriority
	}

	todo := NewTodo(title, description, priority)
	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Service) GetTodo(ctx context.Context, id uuid.UUID) (*Todo, error) {
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, ErrTodoNotFound
	}
	return todo, nil
}

func (s *Service) GetAllTodos(ctx context.Context) ([]*Todo, int, error) {
	todos, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return todos, count, nil
}

func (s *Service) UpdateTodo(ctx context.Context, id uuid.UUID, title *string, description **string, completed *bool, priority *Priority) (*Todo, error) {
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, ErrTodoNotFound
	}

	if title != nil {
		if err := s.validateTitle(*title); err != nil {
			return nil, err
		}
		todo.UpdateTitle(*title)
	}

	if description != nil {
		if err := s.validateDescription(*description); err != nil {
			return nil, err
		}
		todo.UpdateDescription(*description)
	}

	if completed != nil {
		if *completed {
			todo.MarkCompleted()
		} else {
			todo.MarkIncomplete()
		}
	}

	if priority != nil {
		if !priority.IsValid() {
			return nil, ErrInvalidPriority
		}
		todo.UpdatePriority(*priority)
	}

	if err := s.repo.Update(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Service) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if todo == nil {
		return ErrTodoNotFound
	}

	return s.repo.Delete(ctx, id)
}

func (s *Service) validateTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return ErrInvalidTitle
	}
	if len(title) > 200 {
		return ErrInvalidTitle
	}
	return nil
}

func (s *Service) validateDescription(description *string) error {
	if description != nil && len(*description) > 1000 {
		return ErrDescriptionTooLong
	}
	return nil
}