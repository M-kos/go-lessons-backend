SELECT SUM( CASE
	            WHEN operation_type = 'credit' THEN -amount
	            WHEN operation_type = 'debit'  THEN amount
	            END ) as balance
FROM operations
WHERE user_id = $1
  AND sequence_number > $2
