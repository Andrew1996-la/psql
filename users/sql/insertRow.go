package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// вставить одну запись
func InsertRow(ctx context.Context, conn *pgx.Conn, name, phone string) error {
	sqlQuery := `
		INSERT INTO users (name, phone)
		VALUES ($1, $2)
	`

	_, err := conn.Exec(ctx, sqlQuery, name, phone)

	return err
}

// вставить за одну операцию множество данных
func FillUsers(ctx context.Context, conn *pgx.Conn) error {
	rows := [][]interface{}{
		{"Ivan", "111"},
		{"Petr", "222"},
		{"Olga", "333"},
		{"Misha", "444"},
		{"Sanya", "555"},
		{"Andrey", "666"},
		{"Alexey", "777"},
		{"Danil", "888"},
		{"Diana", "999"},
		{"Kristina", "1010"},
	}

	_, err := conn.CopyFrom(
		ctx,
		pgx.Identifier{"sql"},
		[]string{"name", "phone"},
		pgx.CopyFromRows(rows),
	)

	return err
}
