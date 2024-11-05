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
   archive BOOLEAN DEFAULT FALSE,
   PRIMARY KEY (task_id)
);
```

