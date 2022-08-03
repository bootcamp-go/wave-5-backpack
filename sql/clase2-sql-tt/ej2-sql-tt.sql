/*-----------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  SQL 2 - MOVIES DB
         Author:	Israel Fabela
	   Language:	mysql  Ver 8.0.29 for macos12.2 on arm64
		  Topic:	Base de Datos - SQL

	© Mercado Libre - IT Bootcamp 2022

-------------------------------------------------------------------------*/

-- 1. Mostrar el título y el nombre del género de todas las series.
SELECT se.title, ge.name 
	FROM movies_db.series se 
		INNER JOIN movies_db.genres ge ON se.genre_id = ge.id;

-- 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT ep.title, ac.first_name, ac.last_name 
	FROM movies_db.episodes ep 
		JOIN movies_db.actor_episode ac_ep ON ep.id = ac_ep.episode_id
		JOIN movies_db.actors ac ON ac_ep.actor_id = ac.id;

-- 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT se.title, MAX(sea.number) as total_temporadas
	FROM movies_db.series se
		JOIN movies_db.seasons sea ON se.id = sea.serie_id
		GROUP BY se.title ORDER BY total_temporadas DESC;

-- 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT ge.name, COUNT(mo.genre_id) as total_peliculas_genero
	FROM movies_db.genres ge 
		JOIN movies_db.movies mo ON ge.id = mo.genre_id
		GROUP BY mo.genre_id
		HAVING total_peliculas_genero >= 3;

-- 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT ac.first_name, ac.last_name
	FROM  movies_db.actors ac,  movies_db.actor_movie ac_mo, movies_db.movies mo
	WHERE ac.id = ac_mo.actor_id
		AND ac_mo.movie_id = mo.id
		AND mo.title LIKE 'La Guerra de las galaxias%';