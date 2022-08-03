select 
	m.title
from movies_db.movies m 
where year(m.release_date) between 2004 and 2008