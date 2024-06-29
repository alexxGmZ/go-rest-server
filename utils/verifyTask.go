package utils

import (
	"database/sql"
	"fmt"
)

type Task struct {
	TaskId int
}

// Verifies if a task exists in the database by its ID. Returns true if the task exists,
// otherwise returns false.
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
