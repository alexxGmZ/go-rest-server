package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_rest/utils"
)

type Task struct {
	TaskID      int    `json:"task_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Deadline    string `json:"deadline"`
	DateAdded   string `json:"date_added"`
}

func GetTasks(c *gin.Context) {
	query := `
		SELECT task_id, description, status, deadline, date_added
		FROM Tasks
		WHERE archive = FALSE
	`

	rows, err := utils.DB.Query(query)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to query tasks"},
		)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.TaskID, &task.Description, &task.Status, &task.Deadline, &task.DateAdded); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Failed to scan task"},
			)
			return
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating through tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
