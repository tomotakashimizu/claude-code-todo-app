# Database Schema

This directory contains the SQLite database schema and migrations for the Todo application.

## Files

- `schema.sql` - Complete database schema with sample data
- `migrations/001_create_todos_table.sql` - Initial migration to create todos table

## Schema Overview

### Tables

#### `todos`
- `id` (TEXT, PRIMARY KEY) - UUID generated automatically
- `title` (TEXT, NOT NULL) - Todo title (max 200 characters)
- `description` (TEXT, NULLABLE) - Optional description (max 1000 characters)
- `completed` (BOOLEAN, DEFAULT FALSE) - Completion status
- `priority` (TEXT, DEFAULT 'medium') - Priority level: 'low', 'medium', 'high'
- `created_at` (DATETIME, DEFAULT CURRENT_TIMESTAMP) - Creation timestamp
- `updated_at` (DATETIME, DEFAULT CURRENT_TIMESTAMP) - Last update timestamp

### Indexes

- `idx_todos_completed` - Index on completion status
- `idx_todos_created_at` - Index on creation date (DESC)
- `idx_todos_priority` - Index on priority level
- `idx_todos_completed_priority` - Composite index on completion and priority

### Triggers

- `update_todos_updated_at` - Automatically updates `updated_at` field on record updates

## Usage

### Initialize Database

```bash
sqlite3 todos.db < database/schema.sql
```

### Apply Migration

```bash
sqlite3 todos.db < database/migrations/001_create_todos_table.sql
```

### Example Queries

```sql
-- Get all incomplete todos by priority
SELECT * FROM todos 
WHERE completed = FALSE 
ORDER BY 
  CASE priority 
    WHEN 'high' THEN 1 
    WHEN 'medium' THEN 2 
    WHEN 'low' THEN 3 
  END,
  created_at DESC;

-- Mark todo as completed
UPDATE todos SET completed = TRUE WHERE id = 'your-todo-id';

-- Get todos created in the last 7 days
SELECT * FROM todos 
WHERE created_at >= datetime('now', '-7 days')
ORDER BY created_at DESC;
```

## Data Validation

The schema includes CHECK constraints to ensure data integrity:

- Title length: 1-200 characters
- Description length: 0-1000 characters (nullable)
- Priority: Must be 'low', 'medium', or 'high'
- Completed: Boolean values only

## Performance

The database is optimized with strategic indexes for common query patterns:

1. Filtering by completion status
2. Sorting by creation date
3. Filtering by priority
4. Combined completion/priority queries