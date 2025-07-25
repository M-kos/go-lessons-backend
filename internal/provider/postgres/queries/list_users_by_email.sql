SELECT *
FROM users
WHERE email = ANY($1)
