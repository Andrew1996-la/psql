package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func MarkTaskCompleted(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
		UPDATE tasks
		SET completed = true
		where id = $1
	`

	_, err := conn.Exec(ctx, sqlQuery, id)

	return err
}
