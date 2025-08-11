INSERT INTO users (email, full_name, address_id)
VALUES ($1, $2, $3)
RETURNING id, create_time;
