select u.id,
       u.email,
       u.full_name,
       u.create_time,
       a.street,
       a.zip,
       ct.name,
       cn.name
from users as u
	     left join addresses as a on u.address_id = a.id
	     left join cities as ct on a.city_id = ct.id
	     left join countries as cn on ct.country_id = cn.id
where u.email = $1
