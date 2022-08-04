select concat(a.first_name, " " ,a.last_name) actor
from movies_db.actors a 
left join movies_db.movies m on m.id = a.favorite_movie_id 
where m.awards > 3;