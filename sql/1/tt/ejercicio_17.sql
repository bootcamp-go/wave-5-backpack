select 
	m.title
from movies_db.movies m 
where m.rating > 3
and m.awards > 1
and year(m.release_date) between 1988 and 2009
order by m.rating asc