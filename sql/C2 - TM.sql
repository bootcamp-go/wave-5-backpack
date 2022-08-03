/* 1 */ 
SELECT se.title titulo, ge.name genero
FROM series se
	JOIN genres ge ON se.genre_id = ge.id;

/* 2 */
SELECT ep.title titulo, ac.first_name nombre_actor, ac.last_name apellido_actor
FROM episodes ep
	JOIN actor_episode ae ON ep.id = ae.episode_id
	JOIN actors ac ON ac.id = ae.actor_id;

/* 3 */
SELECT title titulo, count(season_id) temporadas
FROM episodes
GROUP BY titulo
ORDER BY temporadas DESC;

/* 4 */
SELECT ge.name genero, count(mo.id) peliculas
FROM genres ge
	JOIN movies mo ON ge.id = mo.genre_id
GROUP BY ge.name
HAVING peliculas >= 3;

/* 5 */
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


