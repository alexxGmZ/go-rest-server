package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go_rest/utils"
)

// Marks a task as done and archives it in the database by its ID.
// Requires the "taskId" parameter in the endpoint. It responds with a success message
// or an error message in case of a failure.
func TaskDone(c *gin.Context) {
	taskIdStr := c.Param("taskId")
	taskIdInt, err := strconv.Atoi(taskIdStr)

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

	sqlQuery := `
		UPDATE Tasks
		SET status = 'Done',
		archive = TRUE
		WHERE task_id = $1
	`

	_, err = utils.DB.Exec(sqlQuery, taskIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to archive task",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task archived successfully"})
}

// Archives a task in the database by its ID. Requires the "taskId" parameter in the
// endpoint. It responds with a success message or an error message in case of a failure.
func ArchiveTask(c *gin.Context) {
	taskIdStr := c.Param("taskId")
	taskIdInt, err := strconv.Atoi(taskIdStr)

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

	sqlQuery := `
		UPDATE Tasks
		SET archive = TRUE
		WHERE task_id = $1
	`

	_, err = utils.DB.Exec(sqlQuery, taskIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to archive task",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task archived successfully"})
}

// Unarchives a task in the database by its ID. Requires the "taskId" parameter in the
// endpoint. It responds with a success message or an error message in case of a failure.
func UnArchiveTask(c *gin.Context) {
	taskIdStr := c.Param("taskId")
	taskIdInt, err := strconv.Atoi(taskIdStr)

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

	sqlQuery := `
		UPDATE Tasks
		SET archive = FALSE
		WHERE task_id = $1
	`

	_, err = utils.DB.Exec(sqlQuery, taskIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to unarchive task",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task unarchived successfully"})
}
