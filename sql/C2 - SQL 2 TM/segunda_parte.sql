#-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
	              #Práctica C2 - SQL 2 TM
		    #Luz Carime Lucumí Hernández
#-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-

-- 1. Mostrar el título y el nombre del género de todas las series.
SELECT se.title, gen.name FROM genres gen 
INNER JOIN series se
ON gen.id = se.genre_id;

-- 2. Mostrar el título de los episodios, el nombre y apellido de los actores 
-- que trabajan en cada uno de ellos.
SELECT ep.title, ac.first_name, ac.last_name 
FROM actors ac
INNER JOIN actor_episode ac_ep ON ac.id = ac_ep.actor_id
INNER JOIN episodes ep ON ac_ep.episode_id = ep.id;

-- 3. Mostrar el título de todas las series y el total de temporadas que tiene 
-- cada una de ellas.
SELECT se.title, COUNT(*) as temporadas FROM series se
INNER JOIN seasons sea 
ON se.id = sea.serie_id
GROUP BY se.title
ORDER BY temporadas;

-- 4. Mostrar el nombre de todos los géneros y la cantidad total de películas 
-- por cada uno, siempre que sea mayor o igual a 3.
SELECT gen.name, COUNT(*) as cantidad_peliculas FROM genres gen
INNER JOIN movies mo
ON gen.id = mo.genre_id
GROUP BY gen.name
HAVING cantidad_peliculas >= 3;

-- 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas 
-- las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT(ac.first_name), ac.last_name
FROM actors ac
INNER JOIN actor_movie ac_mo ON ac.id = ac_mo.actor_id
INNER JOIN movies mo ON ac_mo.movie_id = mo.id
WHERE mo.title LIKE 'La Guerra de las galaxias%';
