-- Mostrar el título y el nombre del género de todas las series.
SELECT se.title, ge.name
FROM movies_db.series se
INNER JOIN movies_db.genres ge ON se.genre_id = ge.id;


-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT ep.title, mo.first_name, mo.last_name
FROM movies_db.episodes ep 
INNER JOIN movies_db.actor_episode ac ON ep.season_id = ac.episode_id
INNER JOIN movies_db.actors mo ON ac.actor_id = mo.id;


-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT  se.title, count(sea.title), max(sea.number)
FROM movies_db.seasons sea
INNER JOIN movies_db.series se ON sea.serie_id = se.id
group by sea.serie_id;


-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT ge.name, count(*) total
FROM movies_db.genres ge
INNER JOIN movies_db.movies mo ON mo.genre_id = ge.id
GROUP BY name HAVING total >= 3;


-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT ac.first_name, ac.last_name
FROM movies_db.actors ac
INNER JOIN movies_db.actor_movie acm ON ac.id = acm.actor_id
INNER JOIN movies_db.movies mo ON acm.movie_id = mo.id
WHERE mo.title LIKE 'La Guerra de las galaxias%'
GROUP BY ac.id 
HAVING count(mo.id) = (SELECT count(*)
FROM movies_db.movies
WHERE title LIKE 'La Guerra de las galaxias%');

