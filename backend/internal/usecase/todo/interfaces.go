package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/tomotakashimizu/claude-code-todo-app/backend/internal/domain/todo"
)

type GetTodosInput struct{}

type GetTodosOutput struct {
	Todos []*todo.Todo `json:"data"`
	Total int          `json:"total"`
}

type CreateTodoInput struct {
	Title       string        `json:"title"`
	Description *string       `json:"description,omitempty"`
	Priority    todo.Priority `json:"priority"`
}

type CreateTodoOutput struct {
	Todo *todo.Todo `json:"todo"`
}

type GetTodoInput struct {
	ID uuid.UUID `json:"id"`
}

type GetTodoOutput struct {
	Todo *todo.Todo `json:"todo"`
}

type UpdateTodoInput struct {
	ID          uuid.UUID      `json:"id"`
	Title       *string        `json:"title,omitempty"`
	Description **string       `json:"description,omitempty"`
	Completed   *bool          `json:"completed,omitempty"`
	Priority    *todo.Priority `json:"priority,omitempty"`
}

type UpdateTodoOutput struct {
	Todo *todo.Todo `json:"todo"`
}

type DeleteTodoInput struct {
	ID uuid.UUID `json:"id"`
}

type DeleteTodoOutput struct{}

type UseCase interface {
	GetTodos(ctx context.Context, input GetTodosInput) (*GetTodosOutput, error)
	CreateTodo(ctx context.Context, input CreateTodoInput) (*CreateTodoOutput, error)
	GetTodo(ctx context.Context, input GetTodoInput) (*GetTodoOutput, error)
	UpdateTodo(ctx context.Context, input UpdateTodoInput) (*UpdateTodoOutput, error)
	DeleteTodo(ctx context.Context, input DeleteTodoInput) (*DeleteTodoOutput, error)
}
