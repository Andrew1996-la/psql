package main

import (
	"context"
	"fmt"
	"psql/features_postgres/simple_connection"
	"psql/features_postgres/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()
	stringConnection := "postgres://postgres:1111@localhost:5432/postgres"

	conn, err := simple_connection.CreateConnection(ctx, stringConnection)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	rows, err := simple_sql.SelectRows(ctx, conn, 200, 0)
	if err != nil {
		panic(err)
	}

	for _, task := range rows {
		if task.ID == 5 {
			task.Title = "Покормить хомяка"
			task.Description = "Отсыпать хомяку порцию корма"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			err := simple_sql.UpdateTask(ctx, conn, task)
			if err != nil {
				panic(err)
			}
			break
		}
	}

	fmt.Println("success")
}
