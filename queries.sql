-- name: ListTodos :many
SELECT id, description, done, created_at
FROM todos
WHERE done = false
ORDER BY id;

-- name: InsertTodo :exec
INSERT INTO todos (description, done)
VALUES (?, ?)
RETURNING *;

-- name: MarkTodoDone :exec
UPDATE todos SET done = true WHERE id = ?;
