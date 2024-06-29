package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_rest/utils"
)

func TaskDone(c *gin.Context) {
	taskId := c.Param("taskId")

	sqlQuery := `
		UPDATE Tasks
		SET status = 'Done',
		archive = TRUE
		WHERE task_id = $1
	`

	_, err := utils.DB.Exec(sqlQuery, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to archive task",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task archived successfully"})
}
