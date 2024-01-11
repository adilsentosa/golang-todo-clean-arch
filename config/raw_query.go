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
)
