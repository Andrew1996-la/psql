package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context, strConn string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, strConn)
}
