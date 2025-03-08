# Go REST Server

Learning Go by building a TODO REST server with PostgreSql.

<br>

## Database

**db name:** go_todo

table/s:

```sql
CREATE TABLE Tasks (
  task_id SERIAL,
  description TEXT NOT NULL,
  status VARCHAR(40) DEFAULT 'On-going',
  deadline TIMESTAMP,
  date_added TIMESTAMP DEFAULT NOW(),
  archive BOOLEAN DEFAULT FALSE,
  PRIMARY KEY (task_id)
);
```

<br>

## Environment Variable/s

* `POSTGRESURL` - database URL.
  - example: `postgresql://user:password@host/db_name`
* `PORT` - Server Port.

<br>

## Endpoints

### GET /task/:taskId

Retrieves a task from the database by its ID, provided it is not archived.
Requires the "**taskId**" parameter in the endpoint. It responds with the task details
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

Counts all active tasks in the database where the deadline is in the future and the
task is not archived. Responds with the task count or an error message in case of a
failure.

**Response:**

StatusOk
```
[Integer]
```

StatusInternalServerError
```json
{
  "message": "Failed to query task",
  "error":   "Error object",
}
```

<br>

### GET /tasks/late

Retrieves a list of tasks from the database where the deadline has already passed.
It responds with a JSON array of late tasks or an error message in case of a failure.

**Response:**

StatusOk
```json
[
  {
    "task_id": 9,
    "description": "love them all",
    "status": "On-going",
    "deadline": "2024-02-14T00:00:00Z",
    "date_added": "2024-06-29T19:03:20.884892Z"
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

### GET /tasks/late/count

Counts all late tasks in the database where the deadline has passed and the task
is not archived. Responds with the count of late tasks or an error message in case of a
failure.

**Response:**

StatusOk
```
[Integer]
```

StatusInternalServerError
```json
{
  "message": "Failed to query task",
  "error":   "Error object",
}
```

<br>

### DELETE /task/:taskId

Deletes a task from the database by its ID. Requires the "**taskId**" parameter in the
endpoint. It responds with a success message or an error message in case of a failure.

**Response:**

StatusOk
```json
{ "message": "Task deleted successfully" }
```

StatusNotFound
```json
{ "message": "Task not found" }
```

StatusInternalServerError
```json
{
  "message": "Failed to convert int to string",
  "error":   "Error object",
}
```
```json
{
  "message": "Failed to delete task",
  "error":   "Error object",
}
```

<br>

### PATCH /task/done/:taskId

Marks a task as done and archives it in the database by its ID.
Requires the "**taskId**" parameter in the endpoint. It responds with a success message
or an error message in case of a failure.

**Response:**

StatusOk
```json
{ "message": "Task archived successfully" }
```

StatusNotFound
```json
{ "message": "Task not found" }
```

StatusInternalServerError
```json
{
  "message": "Failed to convert int to string",
  "error":   "Error object",
}
```
```json
{
  "message": "Failed to archive task",
  "error":   "Error object",
}
```

<br>

### PATCH /task/archive/:taskId

Archives a task in the database by its ID. Requires the "**taskId**" parameter in the
endpoint. It responds with a success message or an error message in case of a failure.

**Response:**

StatusOk
```json
{ "message": "Task archived successfully" }
```

StatusNotFound
```json
{ "message": "Task not found" }
```

StatusInternalServerError
```json
{
  "message": "Failed to convert int to string",
  "error":   "Error object",
}
```
```json
{
  "message": "Failed to archive task",
  "error":   "Error object",
}
```

<br>

### PATCH /task/unarchive/:taskId

Unarchives a task in the database by its ID. Requires the "**taskId**" parameter in the
endpoint. It responds with a success message or an error message in case of a failure.

**Response:**

StatusOk
```json
{ "message": "Task unarchived successfully" }
```

StatusNotFound
```json
{ "message": "Task not found" }
```

StatusInternalServerError
```json
{
  "message": "Failed to convert int to string",
  "error":   "Error object",
}
```
```json
{
  "message": "Failed to unarchive task",
  "error":   "Error object",
}
```

<br>

### POST /create

Creates a new task in the database. It responds with a success message or an error message
in case of a failure.

**Request:**

```json
{
  "description": "",
  "deadline": ""
}
```

**Response:**

StatusOk
```json
{ "message": "Task created successfully" }
```

StatusInternalServerError
```json
{
  "message": "Failed to bind JSON",
  "error":   "Error object",
}
```
```json
{
  "message": "Failed to create task",
  "error":   "Error object",
}
```

<br>

### PUT /task/update

Updates a specific task in the database based on the provided JSON request format.
The task to update is identified implicitly from the context of the request body.
Responds with a success message or an error message in case of a failure.

**Request:**

```json
{
  "task_id": 10,
  "description": "new year",
  "deadline": "2024-01-01"
}
```

**Response:**

StatusOk
```json
{ "message": "Task updated successfully" }
```

StatusNotFound
```json
{ "message": "Task not found" }
```

StatusInternalServerError
```json
{
  "message": "Failed to bind JSON",
  "error":   "Error object",
}
```
```json
{
  "message": "Failed to update task",
  "error":   "Error object",
}
```
<br>

