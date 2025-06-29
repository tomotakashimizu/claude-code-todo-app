package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/tomotakashimizu/claude-code-todo-app/backend/internal/usecase/todo"
)

type TodoHandler struct {
	todoUseCase todo.UseCase
}

func NewTodoHandler(todoUseCase todo.UseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	output, err := h.todoUseCase.GetTodos(ctx, todo.GetTodosInput{})
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, output)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input todo.CreateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body")
		return
	}

	output, err := h.todoUseCase.CreateTodo(ctx, input)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusCreated, output)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "MISSING_ID", "Todo ID is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "INVALID_ID", "Invalid todo ID format")
		return
	}

	output, err := h.todoUseCase.GetTodo(ctx, todo.GetTodoInput{ID: id})
	if err != nil {
		h.writeErrorResponse(w, http.StatusNotFound, "TODO_NOT_FOUND", err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, output)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "MISSING_ID", "Todo ID is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "INVALID_ID", "Invalid todo ID format")
		return
	}

	var input todo.UpdateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body")
		return
	}

	input.ID = id

	output, err := h.todoUseCase.UpdateTodo(ctx, input)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, output)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "MISSING_ID", "Todo ID is required")
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "INVALID_ID", "Invalid todo ID format")
		return
	}

	_, err = h.todoUseCase.DeleteTodo(ctx, todo.DeleteTodoInput{ID: id})
	if err != nil {
		h.writeErrorResponse(w, http.StatusNotFound, "TODO_NOT_FOUND", err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TodoHandler) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *TodoHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, errorType, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	errorResponse := map[string]string{
		"error":   errorType,
		"message": message,
	}
	
	json.NewEncoder(w).Encode(errorResponse)
}