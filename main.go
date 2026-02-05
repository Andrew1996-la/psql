package main

import (
	"context"
	"fmt"
	"psql/features_postgres/simple_connection"
	"psql/features_postgres/simple_sql"
	"psql/features_postgres/utils"
)

func main() {
	ctx := context.Background()
	stringConnection := "postgres://postgres:1111@localhost:5432/postgres"

	conn, err := simple_connection.CreateConnection(ctx, stringConnection)
	if err != nil {
		panic(err)
	}

	//rows, err := users.Pagination(ctx, conn, 2, 5)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(rows)

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	//if err := simple_sql.InsertRow(
	//	ctx,
	//	conn,
	//	"Покушац",
	//	"Покушать собранный обед",
	//	false,
	//	time.Now(),
	//); err != nil {
	//	panic(err)
	//}

	//if err := simple_sql.MarkTaskCompleted(ctx, conn, 2); err != nil {
	//	panic(err)
	//}

	//if err := simple_sql.DeleteRow(ctx, conn, 2); err != nil {
	//	panic(err)
	//}

	rows, err := simple_sql.SelectRows(ctx, conn, 200, 0)
	if err != nil {
		panic(err)
	}

	utils.PrintTask(rows)

	fmt.Println("success")
}
