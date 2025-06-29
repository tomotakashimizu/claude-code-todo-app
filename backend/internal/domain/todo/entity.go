package todo

import (
	"time"

	"github.com/google/uuid"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Completed   bool      `json:"completed"`
	Priority    Priority  `json:"priority"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTodo(title string, description *string, priority Priority) *Todo {
	now := time.Now()
	return &Todo{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    priority,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Todo) MarkCompleted() {
	t.Completed = true
	t.UpdatedAt = time.Now()
}

func (t *Todo) MarkIncomplete() {
	t.Completed = false
	t.UpdatedAt = time.Now()
}

func (t *Todo) UpdateTitle(title string) {
	t.Title = title
	t.UpdatedAt = time.Now()
}

func (t *Todo) UpdateDescription(description *string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Todo) UpdatePriority(priority Priority) {
	t.Priority = priority
	t.UpdatedAt = time.Now()
}

func (p Priority) IsValid() bool {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return true
	default:
		return false
	}
}