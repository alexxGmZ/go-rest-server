# Go REST Server

Learning Go by building a TODO REST server with PostgreSql.

<br>

## Endpoints

### GET /task/:taskId

Retrieves a task from the database by its ID, provided it is not archived.
Requires the "taskId" parameter in the endpoint. It responds with the task details
in JSON format or an error message if the task is not found or if there is a failure
in querying the database.

**Response:**

StatusOk
```json
{
  "task_id": 8,
  "description": "Feed the pets",
  "status": "On-going",
  "deadline": "2024-09-08T00:00:00Z",
  "date_added": "2024-06-29T19:03:00.226263Z"
}
```

StatusNotFound
```json
{ "message": "Task not found" }
```

StatusInternalServerError
```json
{
  "message": "Failed to query task",
  "error":   "Error object",
}
```

<br>

### GET /tasks/all

Retrieves a list of tasks from the database where the deadline is in the future.
It responds with a JSON array of tasks or an error message in case of a failure.

**Response:**

StatusOk
```json
[
  {
    "task_id": 6,
    "description": "Another Task",
    "status": "On-going",
    "deadline": "2024-12-25T00:00:00Z",
    "date_added": "2024-06-29T12:54:54.404224Z"
  },
  {
    "task_id": 7,
    "description": "Another Task",
    "status": "On-going",
    "deadline": "2024-12-25T00:00:00Z",
    "date_added": "2024-06-29T19:02:02.769241Z"
  },
  {
    "task_id": 8,
    "description": "Feed the pets",
    "status": "On-going",
    "deadline": "2024-09-08T00:00:00Z",
    "date_added": "2024-06-29T19:03:00.226263Z"
  }
]
```

StatusInternalServerError
```json
{
  "message": "Failed to query task",
  "error":   "Error object",
}
```

<br>

### GET /tasks/archived

Retrieves a list of archived tasks from the database. It responds with a JSON array
of archived tasks or an error message in case of a failure.

**Response:**

StatusOk
```json
[
  {
    "task_id": 1,
    "description": "Task 1",
    "status": "Done",
    "deadline": "2024-06-24T12:00:00Z",
    "date_added": "2024-06-14T16:27:25.941532Z"
  },
  {
    "task_id": 2,
    "description": "Task 2",
    "status": "Done",
    "deadline": "2024-06-26T12:00:00Z",
    "date_added": "2024-06-14T16:27:25.941532Z"
  },
  {
    "task_id": 3,
    "description": "Task 3",
    "status": "Done",
    "deadline": "2024-12-25T00:00:00Z",
    "date_added": "2024-06-23T14:43:56.601526Z"
  }
]
```

StatusInternalServerError
```json
{
  "message": "Failed to query task",
  "error":   "Error object",
}
```

<br>

### GET /tasks/count

<br>

### GET /tasks/late

<br>

### GET /tasks/late/count

<br>

### DELETE /task/:taskId

<br>

### PATCH /task/done/:taskId

<br>

### PATCH /task/archive/:taskId

<br>

### PATCH /task/unarchive/:taskId

<br>

### POST /create

<br>

### PUT /task/update

<br>

