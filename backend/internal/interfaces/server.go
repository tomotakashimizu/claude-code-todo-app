package interfaces

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/tomotakashimizu/claude-code-todo-app/backend/internal/domain/todo"
	"github.com/tomotakashimizu/claude-code-todo-app/backend/internal/interfaces/handler"
	"github.com/tomotakashimizu/claude-code-todo-app/backend/internal/infrastructure/database/sqlite"
	todoUseCase "github.com/tomotakashimizu/claude-code-todo-app/backend/internal/usecase/todo"
)

type Server struct {
	httpServer *http.Server
	db         *sql.DB
}

func NewServer(port int, dbPath string) (*Server, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := enableForeignKeys(db); err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	todoRepo := sqlite.NewTodoRepository(db)
	todoService := todo.NewService(todoRepo)
	todoInteractor := todoUseCase.NewInteractor(todoService)
	todoHandler := handler.NewTodoHandler(todoInteractor)

	mux := http.NewServeMux()
	setupRoutes(mux, todoHandler)

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      corsMiddleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		db:         db,
	}, nil
}

func (s *Server) Start() error {
	fmt.Printf("Starting server on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return s.db.Close()
}

func setupRoutes(mux *http.ServeMux, todoHandler *handler.TodoHandler) {
	mux.HandleFunc("GET /api/todos", todoHandler.GetTodos)
	mux.HandleFunc("POST /api/todos", todoHandler.CreateTodo)
	mux.HandleFunc("GET /api/todos/{id}", todoHandler.GetTodo)
	mux.HandleFunc("PUT /api/todos/{id}", todoHandler.UpdateTodo)
	mux.HandleFunc("DELETE /api/todos/{id}", todoHandler.DeleteTodo)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func enableForeignKeys(db *sql.DB) error {
	_, err := db.Exec("PRAGMA foreign_keys = ON")
	return err
}