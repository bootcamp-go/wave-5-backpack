select 
	m.title,
	m.rating 
from movies_db.movies m 
where lower(m.title) like '%toy story%'