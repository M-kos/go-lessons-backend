INSERT INTO cities (name, country_id)
VALUES ($1, $2)
RETURNING id;
