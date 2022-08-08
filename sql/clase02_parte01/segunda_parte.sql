-- Mostrar el título y el nombre del género de todas las series.
	SELECT movies_db.s.title titulo, g.name genero
    FROM series s
    INNER JOIN genres g
    ON s.genre_id = g.id;
	
-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
	SELECT ep.title, ac.first_name, ac.last_name
	FROM movies_db.actor_episode ae, movies_db.actors ac, movies_db.episodes ep
	WHERE ae.actor_id = ac.id
	AND ae.episode_id = ep.id ;
	

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
	SELECT	s.title, count(*) as number_of_seasons
	FROM movies_db.series s
    INNER JOIN movies_db.seasons on s.id = seasons.serie_id
    GROUP BY title;
    
    -- otra forma de hacerlo sería buscando el valor más alto en la columna "number" de la tabla season
    SELECT series.title, MAX(seasons.number) temporadas
	FROM movies_db.series, movies_db.seasons
	WHERE series.id = seasons.serie_id
	GROUP BY series.title ;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
    SELECT	gr.name as genero, count(*) as cant_pelis
	FROM movies_db.genres gr
	INNER JOIN movies_db.movies mo ON gr.id = mo.genre_id
	GROUP BY name
	HAVING cant_pelis >= 3;
    

-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
	SELECT DISTINCT ac.first_name nombre, ac.last_name apellido
	FROM movies_db.actors ac
	INNER JOIN movies_db.actor_movie am ON ac.id = am.actor_id
	WHERE am.movie_id IN (SELECT id FROM movies WHERE title LIKE '%la guerra de las galaxias%');

