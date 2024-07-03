package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go_rest/utils"
)

type Task struct {
	TaskID      int       `json:"task_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Deadline    time.Time `json:"deadline"`
	DateAdded   time.Time `json:"date_added"`
}

// Retrieves a list of tasks from the database where the deadline is in the future.
// It responds with a JSON array of tasks or an error message in case of a failure.
func GetTasks(c *gin.Context) {
	sqlQuery := `
		SELECT task_id, description, status, deadline, date_added
		FROM Tasks
		WHERE deadline > NOW()
		AND archive = FALSE
	`

	tasks, err := queryTasks(sqlQuery)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to query tasks",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Retrieves a list of tasks from the database where the deadline has already passed.
// It responds with a JSON array of late tasks or an error message in case of a failure.
func GetLateTasks(c *gin.Context) {
	sqlQuery := `
		SELECT task_id, description, status, deadline, date_added
		FROM Tasks
		WHERE deadline < NOW()
		AND archive = FALSE
	`

	lateTasks, err := queryTasks(sqlQuery)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to query tasks",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, lateTasks)
}

// Retrieves a task from the database by its ID, provided it is not archived.
// Requires the "taskId" parameter in the endpoint. It responds with the task details
// in JSON format or an error message if the task is not found or if there is a failure
// in querying the database.
func GetTaskById(c *gin.Context) {
	taskId := c.Param("taskId")
	sqlQuery := `
		SELECT task_id, description, status, deadline, date_added
		FROM Tasks
		WHERE task_id = $1
	`

	row := utils.DB.QueryRow(sqlQuery, taskId)

	var task Task
	err := row.Scan(
		&task.TaskID,
		&task.Description,
		&task.Status,
		&task.Deadline,
		&task.DateAdded,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to query task",
				"error":   err,
			})
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

// Retrieves a list of archived tasks from the database. It responds with a JSON array
// of archived tasks or an error message in case of a failure.
func GetArchivedTasks(c *gin.Context) {
	sqlQuery := `
      SELECT task_id, description, status, deadline, date_added
      FROM Tasks
      WHERE archive = TRUE
   `

	tasks, err := queryTasks(sqlQuery)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to query task",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Executes a given SQL query to retrieve tasks from the database.
// Returns a slice of Task objects and an error if the query fails or if there
// is an issue iterating through the result set.
func queryTasks(sqlQuery string) ([]Task, error) {
	rows, err := utils.DB.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(
			&task.TaskID,
			&task.Description,
			&task.Status,
			&task.Deadline,
			&task.DateAdded,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
