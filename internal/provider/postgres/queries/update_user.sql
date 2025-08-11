UPDATE users
SET email=$1,
    full_name=$2,
    address_id=$3
WHERE id = $4
