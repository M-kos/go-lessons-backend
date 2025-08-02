select id, name
from cities
where name = $1
  and country_id = $2
