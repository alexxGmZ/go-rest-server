package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_rest/utils"
)

type NewTask struct {
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

// Creates a new task in the database. Requires a JSON request with the following format:
// {
//   "Description": "",
//   "Deadline": ""
// }
// It responds with a success message or an error message in case of a failure.
func CreateTask(c *gin.Context) {
	var newTask NewTask

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to bind JSON",
			"error":   err,
		})
		return
	}

	sqlQuery := `
		INSERT INTO Tasks (description, deadline)
		VALUES ($1, $2)
	`

	_, err := utils.DB.Exec(sqlQuery, newTask.Description, newTask.Deadline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create task",
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "message": "Task created successfully" })
}
