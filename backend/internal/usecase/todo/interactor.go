package todo

import (
	"context"

	"github.com/todo-app/backend/internal/domain/todo"
)

type Interactor struct {
	todoService *todo.Service
}

func NewInteractor(todoService *todo.Service) *Interactor {
	return &Interactor{
		todoService: todoService,
	}
}

func (i *Interactor) GetTodos(ctx context.Context, input GetTodosInput) (*GetTodosOutput, error) {
	todos, total, err := i.todoService.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}

	return &GetTodosOutput{
		Todos: todos,
		Total: total,
	}, nil
}

func (i *Interactor) CreateTodo(ctx context.Context, input CreateTodoInput) (*CreateTodoOutput, error) {
	createdTodo, err := i.todoService.CreateTodo(ctx, input.Title, input.Description, input.Priority)
	if err != nil {
		return nil, err
	}

	return &CreateTodoOutput{
		Todo: createdTodo,
	}, nil
}

func (i *Interactor) GetTodo(ctx context.Context, input GetTodoInput) (*GetTodoOutput, error) {
	foundTodo, err := i.todoService.GetTodo(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &GetTodoOutput{
		Todo: foundTodo,
	}, nil
}

func (i *Interactor) UpdateTodo(ctx context.Context, input UpdateTodoInput) (*UpdateTodoOutput, error) {
	updatedTodo, err := i.todoService.UpdateTodo(
		ctx,
		input.ID,
		input.Title,
		input.Description,
		input.Completed,
		input.Priority,
	)
	if err != nil {
		return nil, err
	}

	return &UpdateTodoOutput{
		Todo: updatedTodo,
	}, nil
}

func (i *Interactor) DeleteTodo(ctx context.Context, input DeleteTodoInput) (*DeleteTodoOutput, error) {
	err := i.todoService.DeleteTodo(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteTodoOutput{}, nil
}