CREATE TABLE IF NOT EXISTS users(
  user_id serial PRIMARY KEY,
  user_name VARCHAR(50) NOT NULL,
  birthday DATE NOT NULL
);