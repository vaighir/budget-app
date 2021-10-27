CREATE TABLE users (
  id serial PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE foods (
  id serial PRIMARY KEY,
  food_name VARCHAR(255) UNIQUE NOT NULL,
  calories_per_gram INTEGER
);
