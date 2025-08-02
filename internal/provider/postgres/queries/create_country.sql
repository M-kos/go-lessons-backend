INSERT INTO countries (name)
VALUES ($1)
RETURNING id;
