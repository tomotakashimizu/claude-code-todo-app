-- Migration: 001_create_todos_table.sql
-- Description: Create the initial todos table with indexes and triggers
-- Created: 2025-06-29

-- Enable foreign key constraints
PRAGMA foreign_keys = ON;

-- Create todos table
CREATE TABLE todos (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    title TEXT NOT NULL CHECK(length(title) <= 200 AND length(title) > 0),
    description TEXT CHECK(description IS NULL OR length(description) <= 1000),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    priority TEXT CHECK(priority IN ('low', 'medium', 'high')) DEFAULT 'medium',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create performance indexes
CREATE INDEX idx_todos_completed ON todos(completed);
CREATE INDEX idx_todos_created_at ON todos(created_at DESC);
CREATE INDEX idx_todos_priority ON todos(priority);
CREATE INDEX idx_todos_completed_priority ON todos(completed, priority);

-- Create trigger for automatic updated_at timestamp
CREATE TRIGGER update_todos_updated_at
    AFTER UPDATE ON todos
    FOR EACH ROW
    WHEN NEW.updated_at = OLD.updated_at
BEGIN
    UPDATE todos SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;