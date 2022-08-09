/*
Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
*/

SELECT a.first_name,a.last_name,m.awards
FROM actors a
LEFT JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards >3