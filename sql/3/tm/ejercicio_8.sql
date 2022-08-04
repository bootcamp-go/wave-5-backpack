select g.name 
from movies_db.genres g
left join movies_db.movies m on m.genre_id = g.id
group by g.id
having count(m.id) > 0;