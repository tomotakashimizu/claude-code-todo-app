-- Todo Application Database Schema
-- SQLite database schema for the Todo application

-- Enable foreign key constraints
PRAGMA foreign_keys = ON;

-- Main todos table
CREATE TABLE todos (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    title TEXT NOT NULL CHECK(length(title) <= 200 AND length(title) > 0),
    description TEXT CHECK(description IS NULL OR length(description) <= 1000),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    priority TEXT CHECK(priority IN ('low', 'medium', 'high')) DEFAULT 'medium',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Performance indexes
CREATE INDEX idx_todos_completed ON todos(completed);
CREATE INDEX idx_todos_created_at ON todos(created_at DESC);
CREATE INDEX idx_todos_priority ON todos(priority);
CREATE INDEX idx_todos_completed_priority ON todos(completed, priority);

-- Trigger to automatically update updated_at timestamp
CREATE TRIGGER update_todos_updated_at
    AFTER UPDATE ON todos
    FOR EACH ROW
    WHEN NEW.updated_at = OLD.updated_at
BEGIN
    UPDATE todos SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- Insert some sample data for development
INSERT INTO todos (title, description, completed, priority) VALUES
    ('Learn OpenAPI', 'Study OpenAPI 3.0 specifications and best practices', FALSE, 'high'),
    ('Setup database', 'Create SQLite schema and migrations for the Todo app', TRUE, 'high'),
    ('Implement API endpoints', 'Create REST API endpoints for CRUD operations', FALSE, 'medium'),
    ('Add frontend components', 'Build React components for todo management', FALSE, 'medium'),
    ('Write tests', 'Add unit and integration tests for the application', FALSE, 'low');