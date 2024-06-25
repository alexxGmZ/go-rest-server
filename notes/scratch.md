Database: postgresql

**db name:** go_todo

table/s:

```sql
CREATE TABLE Tasks (
   task_id SERIAL,
   description TEXT NOT NULL,
   status VARCHAR(40) DEFAULT 'On-going',
   deadline TIMESTAMP,
   date_added TIMESTAMP DEFAULT NOW(),
   archive BOOLEAN DEFAULT TRUE,
   PRIMARY KEY (task_id)
);
```

## Planned endpoints

* [x] GET /tasks
* [x] GET /tasks/:taskId
* [ ] GET /archived
* [ ] GET /late
* [ ] GET /late/count
* [ ] GET /count
* [ ] PUT /update/:taskId
* [ ] PATCH /archive/:taskId
* [x] POST /task
* [x] DELETE /task/:taskId
