INSERT INTO users (email, full_name)
VALUES ($1, $2)
RETURNING id, create_time;
