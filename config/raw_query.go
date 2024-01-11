package config

const (
	SelectAuthorById      = "SELECT id,name,email,created_at,updated_at FROM authors WHERE id = $1"
	SelectAuthorByEmail   = "SELECT id,name,email,created_at, updated_at FROM authors WHERE email = $1"
	SelectAuthorWithTasks = `
  SELECT
	  a.id,
	  a.name,
	  a.email,
    a.role,
    a.updated_at,
	  a.created_at,
	  t.id  as t_id,
	  t.title as t_title,
	  t.content as t_content,
    t.author_id as t_author_id,
	  t.created_at as t_created_at,
	  t.updated_at as t_updated_at
  FROM
	  authors a
  JOIN
	  tasks t  ON a.id = t.author_id
  ORDER BY a.email`

	SelectAuthorWithTasksByID = `
  SELECT
	  a.id,
	  a.name,
	  a.email,
    a.role,
    a.updated_at,
	  a.created_at,
	  t.id  as t_id,
	  t.title as t_title,
	  t.content as t_content,
    t.author_id as t_author_id,
	  t.created_at as t_created_at,
	  t.updated_at as t_updated_at
  FROM
	  authors a
  JOIN
	  tasks t  ON a.id = t.author_id
  WHERE a.id = $1`

	InsertIntoTask = "INSERT INTO tasks (title, content,author_id,updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at"

	SelectTaskByAuthorID = "SELECT id,title,content,created_at,updated_at FROM tasks WHERE author_id = $1"

	SelectTaskPagination = `SELECT id,title,content,author_id,created_at 
  FROM tasks ORDER BY created_at DESC LIMIT $1 OFFSET $2`
)
