CREATE TABLE role (
    id serial PRIMARY KEY,
    name VARCHAR(255)
);


CREATE TABLE users (
  id serial PRIMARY KEY,
  role_id INT NOT NULL DEFAULT 2,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  CONSTRAINT role_id
    FOREIGN KEY (role_id)
    REFERENCES role(id)
    ON DELETE CASCADE
);

CREATE TABLE foods (
  id serial PRIMARY KEY,
  food_name VARCHAR(255) UNIQUE NOT NULL,
  calories_per_gram INTEGER
);
