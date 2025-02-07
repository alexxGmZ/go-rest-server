package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go_rest/utils"
)

// Deletes a task from the database by its ID. Requires the "taskId" parameter in the
// endpoint. It responds with a success message or an error message in case of a failure.
func DeleteTask(c *gin.Context) {
	taskIdStr := c.Param("taskId")
	taskIdInt, err := strconv.Atoi(taskIdStr)

	sqlQuery := `
		DELETE FROM Tasks
		WHERE task_id = $1
	`

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to convert int to string",
			"error":   err,
		})
		return
	}

	if !utils.VerifyTask(taskIdInt) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	_, err = utils.DB.Exec(sqlQuery, taskIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete task",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
