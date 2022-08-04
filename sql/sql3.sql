
Mostrar el título y el nombre del género de todas las series.

SELECT mo.title as titulo,
 gn.name as genero 
 FROM movies mo 
 INNER JOIN genres gn;

Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

SELECT ep.title titulo, ac.first_name nombre_actor, ac.last_name apellido_actor
FROM episodes ep
	JOIN actor_episode ae ON ep.id = ae.episode_id
	JOIN actors ac ON ac.id = ae.actor_id;

Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

SELECT title titulo, count(season_id) temporadas
FROM episodes
GROUP BY titulo
ORDER BY temporadas DESC;

Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

SELECT ge.name genero, count(mo.id) peliculas
FROM genres ge
	JOIN movies mo ON ge.id = mo.genre_id
GROUP BY ge.name
HAVING peliculas >= 3;

Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.

SELECT ac.first_name nombre, ac.last_name apellido
FROM actors ac
	JOIN actor_movie am ON ac.id = am.actor_id
	JOIN movies mo ON mo.id = am.movie_id
WHERE mo.title LIKE "La Guerra de las galaxias%"
GROUP BY ac.id
HAVING count(mo.id) = (
	SELECT count(*)
    FROM movies
    WHERE title LIKE "La Guerra de las galaxias%"
);