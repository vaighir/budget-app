CREATE TABLE role (
    id serial PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE household (
  id serial PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  months_for_emergency_fund INT
);

CREATE TABLE users (
  id serial PRIMARY KEY,
  role_id INT NOT NULL DEFAULT 1,
  household_id INT NOT NULL,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  CONSTRAINT role_id
    FOREIGN KEY (role_id)
    REFERENCES role(id)
    ON DELETE CASCADE,
  CONSTRAINT household_id
    FOREIGN KEY (household_id)
    REFERENCES household(id)
    ON DELETE CASCADE
);


CREATE TABLE savings (
  id serial PRIMARY KEY,
  household_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  amount FLOAT NOT NULL,
  CONSTRAINT household_id
    FOREIGN KEY (household_id)
    REFERENCES household(id)
    ON DELETE CASCADE
);

CREATE TABLE income (
  id serial PRIMARY KEY,
  household_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  amount FLOAT NOT NULL,
  CONSTRAINT household_id
    FOREIGN KEY (household_id)
    REFERENCES household(id)
    ON DELETE CASCADE
);

CREATE TABLE monthly_expences (
  id serial PRIMARY KEY,
  household_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  amount FLOAT NOT NULL,
  CONSTRAINT household_id
    FOREIGN KEY (household_id)
    REFERENCES household(id)
    ON DELETE CASCADE
);

CREATE TABLE upcoming_expences (
  id serial PRIMARY KEY,
  household_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  amount FLOAT NOT NULL,
  deadline DATE NOT NULL,
  CONSTRAINT household_id
    FOREIGN KEY (household_id)
    REFERENCES household(id)
    ON DELETE CASCADE
);

CREATE TABLE funds (
  id serial PRIMARY KEY,
  household_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  amount FLOAT NOT NULL,
  CONSTRAINT household_id
    FOREIGN KEY (household_id)
    REFERENCES household(id)
    ON DELETE CASCADE
);
