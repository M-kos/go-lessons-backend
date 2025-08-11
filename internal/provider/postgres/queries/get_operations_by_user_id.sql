SELECT id, user_id, sequence_number, operation_type, amount, create_time
FROM operations
WHERE user_id = $1
