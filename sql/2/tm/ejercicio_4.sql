select 
	g.name,
	count(m.id) movies
from movies_db.genres g 
join movies_db.movies m on m.genre_id = g.id 
group by g.id 
having count(m.id) >= 3