package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	stringConneciton:= "postgres://postgres:1111@localhost:5432/postgres"

	conn, err := pgx.Connect(ctx, stringConneciton)
	if err != nil {
		panic(err)
	}

	if err:=conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Успешное подключение к базе данных")
}
