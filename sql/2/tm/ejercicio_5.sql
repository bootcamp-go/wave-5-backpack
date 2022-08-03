select distinct
	a.first_name,
	a.last_name 
from movies_db.actors a 
join movies_db.actor_movie am on am.actor_id = a.id 
join movies_db.movies m on m.id = am.movie_id
where lower(m.title) like '%la guerra de las galaxias%'