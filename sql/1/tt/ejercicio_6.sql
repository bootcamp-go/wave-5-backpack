select 
	m.title,
	m.rating,
	m.awards 
from movies_db.movies m
where m.rating > 7.5
and m.awards > 2