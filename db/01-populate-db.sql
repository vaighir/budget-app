INSERT INTO role (name) values ('admin');
INSERT INTO role (name) values ('user');

INSERT INTO household (name) values ('admin');

INSERT INTO users (role_id, household_id, username, password) values (1, 1, 'admin', '$2a$10$Q8Q4T/PQf2tJ0q.lnMdlNe9ndpIfT7fvOIGgg.0nTwE0qF9l.qQsi');
