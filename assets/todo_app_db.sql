CREATE DATABASE todo_app_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE authors (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name VARCHAR(50),
  email VARCHAR(100) UNIQUE,
  password VARCHAR(100),
  role VARCHAR(50),
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp
)

CREATE TABLE tasks (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  title VARCHAR(100),
  content TEXT,
  author_id uuid,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp
  CONSTRAINT task_author_id_fkey FOREIGN KEY (author_id)
    REFERENCES authors(id)
)

