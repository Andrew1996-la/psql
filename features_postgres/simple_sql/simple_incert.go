package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		INSERT INTO tasks(title, description, completed, created_at)
		VALUES ('Домашка', 'Сделать домашку по математике', FALSE, '2026.01.02 14:30:05') 
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
