package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func SelectRows(
	ctx context.Context,
	conn *pgx.Conn,
	limit, offset int,
) ([]Task, error) {
	sqlQuery := `
			SELECT id, title, description, completed, created_at, completed_at
			FROM tasks
			ORDER BY id
			LIMIT $1
			OFFSET $2
		`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.CompletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
