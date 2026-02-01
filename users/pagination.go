package users

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Name  string
	Phone string
}

func Pagination(
	ctx context.Context,
	conn *pgx.Conn,
	offset, limit int,
) ([]User, error) {
	sqlQuery := `
		SELECT name, phone FROM users
		ORDER BY id
		LIMIT $1
		OFFSET $2
	`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Name, &user.Phone); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
