package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_rest/utils"
)

// Deletes a task from the database by its ID. Requires the "taskId" parameter in the
// endpoint. It responds with a success message or an error message in case of a failure.
func DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")

	sqlQuery := `
		DELETE FROM Tasks
		WHERE task_id = $1
	`

	_, err := utils.DB.Exec(sqlQuery, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete task",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
