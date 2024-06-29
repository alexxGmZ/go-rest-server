package utils

import (
	"database/sql"
	"fmt"
)

type Task struct {
	TaskId int
}

func VerifyTask(taskId int) bool {
	fmt.Printf("VerifyTask(%d)\n", taskId)

	sqlQuery := `
		SELECT task_id FROM Tasks
		WHERE task_id = $1
	`
	row := DB.QueryRow(sqlQuery, taskId)
	var task Task
	err := row.Scan(&task.TaskId)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	return true
}
