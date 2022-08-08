USE movies_db;

SELECT ac.first_name nombre, ac.last_name apellido,mo.title pelicula
FROM actors ac
INNER JOIN movies mo ON mo.id = ac.favorite_movie_id
WHERE mo.awards > 3;