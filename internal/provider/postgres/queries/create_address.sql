INSERT INTO addresses (street, zip, city_id)
VALUES ($1, $2, $3)
RETURNING id;
