select 
	s.title,
	g.name 
from movies_db.series s 
join movies_db.genres g on g.id = s.genre_id 