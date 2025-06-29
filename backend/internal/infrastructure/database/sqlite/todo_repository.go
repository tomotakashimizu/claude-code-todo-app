package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomotakashimizu/claude-code-todo-app/backend/internal/domain/todo"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Create(ctx context.Context, t *todo.Todo) error {
	query := `
		INSERT INTO todos (id, title, description, completed, priority, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	
	_, err := r.db.ExecContext(
		ctx,
		query,
		t.ID.String(),
		t.Title,
		t.Description,
		t.Completed,
		string(t.Priority),
		t.CreatedAt,
		t.UpdatedAt,
	)
	
	return err
}

func (r *TodoRepository) GetByID(ctx context.Context, id uuid.UUID) (*todo.Todo, error) {
	query := `
		SELECT id, title, description, completed, priority, created_at, updated_at
		FROM todos
		WHERE id = ?
	`
	
	row := r.db.QueryRowContext(ctx, query, id.String())
	
	var t todo.Todo
	var idStr string
	var description sql.NullString
	var priorityStr string
	
	err := row.Scan(
		&idStr,
		&t.Title,
		&description,
		&t.Completed,
		&priorityStr,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	
	t.ID, err = uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}
	
	if description.Valid {
		t.Description = &description.String
	}
	
	t.Priority = todo.Priority(priorityStr)
	
	return &t, nil
}

func (r *TodoRepository) GetAll(ctx context.Context) ([]*todo.Todo, error) {
	query := `
		SELECT id, title, description, completed, priority, created_at, updated_at
		FROM todos
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var todos []*todo.Todo
	
	for rows.Next() {
		var t todo.Todo
		var idStr string
		var description sql.NullString
		var priorityStr string
		
		err := rows.Scan(
			&idStr,
			&t.Title,
			&description,
			&t.Completed,
			&priorityStr,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		
		t.ID, err = uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}
		
		if description.Valid {
			t.Description = &description.String
		}
		
		t.Priority = todo.Priority(priorityStr)
		
		todos = append(todos, &t)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return todos, nil
}

func (r *TodoRepository) Update(ctx context.Context, t *todo.Todo) error {
	query := `
		UPDATE todos
		SET title = ?, description = ?, completed = ?, priority = ?, updated_at = ?
		WHERE id = ?
	`
	
	result, err := r.db.ExecContext(
		ctx,
		query,
		t.Title,
		t.Description,
		t.Completed,
		string(t.Priority),
		time.Now(),
		t.ID.String(),
	)
	
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return todo.ErrTodoNotFound
	}
	
	return nil
}

func (r *TodoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM todos WHERE id = ?`
	
	result, err := r.db.ExecContext(ctx, query, id.String())
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return todo.ErrTodoNotFound
	}
	
	return nil
}

func (r *TodoRepository) Count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM todos`
	
	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	
	return count, nil
}