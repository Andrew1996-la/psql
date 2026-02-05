package utils

import (
	"fmt"
	"psql/features_postgres/simple_sql"
)

func PrintTask(tasks []simple_sql.TaskModel) {
	for _, task := range tasks {
		fmt.Println("-------------------------------------")
		fmt.Println("id:", task.ID)
		fmt.Println("title:", task.Title)
		fmt.Println("description:", task.Description)
		fmt.Println("completed:", task.Completed)
		fmt.Println("created_at:", task.CreatedAt)
		fmt.Println("completed_at:", task.CompletedAt)
	}
}
