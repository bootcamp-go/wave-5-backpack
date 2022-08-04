USE movies_db;

/* EJERCICIO 1 SEGUNDA PARTE*/
SELECT serie.title as titulo, gen.name as genero
FROM series as serie
INNER JOIN genres as gen 
WHERE serie.genre_id = gen.id; 

/* EJERCICIO 2 SEGUNDA PARTE*/
SELECT
  episodes.title as titulo,
  actors.first_name as nombre,
  actors.last_name as apellido
FROM episodes
JOIN actor_episode
  ON episodes.id = actor_episode.episode_id
JOIN actors
  ON actor_episode.actor_id = actors.id;
  
/* EJERCICIO 3 SEGUNDA PARTE*/
SELECT
  series.title as titulo,
  count(*)
FROM series
JOIN seasons
  ON series.id=seasons.serie_id
GROUP BY
	series.title;

/* EJERCICIO 4 SEGUNDA PARTE*/
SELECT
  genres.`name` as genero,
  count(*) as numeritos
FROM genres
JOIN movies
  ON genres.id=movies.genre_id
GROUP BY
	genres.`name`
HAVING
	numeritos>3;

/* EJERCICIO 5 SEGUNDA PARTE*/
SELECT actors.first_name, actors.last_name FROM actors
 WHERE actors.id IN
	(SELECT
	  actors.id
	FROM actors
	JOIN actor_movie
	  ON actors.id=actor_movie.actor_id
	JOIN movies
	  ON actor_movie.movie_id = movies.id
	WHERE 
		movies.id IN (SELECT movies.id WHERE movies.title like "%guerra de las galaxias%") ## no se porque funciona esta subquery
	)
GROUP BY 
	actors.id;
/*
Error Code: 1055. Expression #2 of SELECT list is not in GROUP BY clause and contains nonaggregated column 'movies_db.actors.last_name' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by


