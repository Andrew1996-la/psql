package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(
	ctx context.Context,
	conn *pgx.Conn,
	limit, offset int,
) ([]TaskModel, error) {
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

	var tasks []TaskModel

	for rows.Next() {
		var task TaskModel
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.CompletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
