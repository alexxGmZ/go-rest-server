package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_rest/utils"
)

type TaskUpdate struct {
	TaskID      int    `json:"task_id"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

// Updates a specific task in the database based on the provided JSON request format.
// Requires the JSON request in the following format:
//
//	{
//	   "task_id": 10,
//	   "description": "new year",
//	   "deadline": "2024-01-01"
//	}
//
// The task to update is identified implicitly from the context of the request body.
// Responds with a success message or an error message in case of a failure.
func UpdateSpecificTask(c *gin.Context) {
	var taskUpdate TaskUpdate

	if err := c.BindJSON(&taskUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to bind JSON",
			"error":   err,
		})
		return
	}

	if !utils.VerifyTask(taskUpdate.TaskID) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	sqlQuery := `
      UPDATE Tasks SET
      description = $1,
      deadline = $2
      WHERE task_id = $3
   `
	_, err := utils.DB.Exec(
		sqlQuery,
		taskUpdate.Description,
		taskUpdate.Deadline,
		taskUpdate.TaskID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update task",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
