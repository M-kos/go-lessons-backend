select id, street, zip
from addresses
where street = $1
  and zip = $2
  and city_id = $3
