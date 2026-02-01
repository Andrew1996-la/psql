package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY,
		    name VARCHAR(255) NOT NULL,
		    phone VARCHAR(255) NOT NULL
		)
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
