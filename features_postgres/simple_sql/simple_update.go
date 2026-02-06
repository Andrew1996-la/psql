package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func UpdateTask(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlQuery := `
		UPDATE tasks
		SET title = $1, description = $2, completed = $3, completed_at=$4
		where id = $5
	`

	ct, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CompletedAt, task.ID)

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("task '%s' not found", task.ID)
	}

	return err
}
