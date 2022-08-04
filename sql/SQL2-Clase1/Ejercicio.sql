-- 1 Mostrar el título y el nombre del género de todas las series.
SELECT se.title as titulo, generos.name as genero
FROM series se
INNER JOIN genres generos
ON se.genre_id=generos.id;
-- 2 Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT ep.title as titulo, actores.first_name as nombre , actores.last_name as apellido
	FROM episodes ep
	INNER JOIN actor_episode act_ep
    ON ep.id=act_ep.episode_id
	INNER JOIN actors actores
	ON act_ep.actor_id=actores.id;
-- 3 Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT se.title as titulo, count(*) as cantidad
FROM series se
	INNER JOIN seasons temporadas
    ON se.id=temporadas.serie_id
    GROUP BY se.id;
-- 4 Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT generos.name as genero, count(*) as cantidad
FROM movies peliculas
	INNER JOIN genres generos
    ON generos.id=peliculas.genre_id
    GROUP BY generos.name HAVING cantidad>=3;
-- 5 Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT actores.first_name as nombre , actores.last_name as apellido
	FROM actors actores
	INNER JOIN actor_movie act_peli
    ON actores.id=act_peli.movie_id
    INNER JOIN movies pelis
    ON act_peli.movie_id=pelis.id
    WHERE pelis.title LIKE 'La Guerra de las galaxias%'
    GROUP BY actores.id;