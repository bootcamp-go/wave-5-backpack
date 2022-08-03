/*Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas 
de la guerra de las galaxias y que estos no se repitan.*/

SELECT DISTINCT a.first_name,a.last_name, m.title
FROM actors a
INNER JOIN movies m ON m.id = a.favorite_movie_id
WHERE m.title LIKE "%La Guerra de las galaxias%"