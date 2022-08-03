select 
	s.title,
	count(s2.id) seasons
from movies_db.series s 
join movies_db.seasons s2 on s2.serie_id = s.id 
group by s.id 
