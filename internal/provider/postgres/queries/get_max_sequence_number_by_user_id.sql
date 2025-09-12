SELECT COALESCE( MAX( sequence_number ), 0 )
FROM operations
WHERE user_id = $1
