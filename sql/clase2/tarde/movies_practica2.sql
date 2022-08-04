USE movies_db;

#1 mostrar titulo y nombre del genero de todas las series
SELECT series.title, genres.name FROM genres INNER JOIN series ON genres.id = series.genre_id;

#2 mostrar titulo episodio, actores nombre, apellido
SELECT actor_episode.actor_id, actor_episode.episode_id, 
actors.first_name AS nameActor, actors.last_name AS apellidoActor, 
episodes.title AS tituloEpisodio
FROM actor_episode INNER JOIN
actors ON actor_episode.actor_id = actors.id 
INNER JOIN episodes ON episodes.id = actor_episode.episode_id;

#3 mostrar titulo todas series y total de temporadas de c/u 
SELECT series.title, COUNT(*) AS NoTemporadas FROM seasons LEFT JOIN series ON seasons.serie_id = series.id 
GROUP BY series.title;

#4 nombre de todos los generos y cantidad total de pelis por cada uno siempre y cuandos sea >= 3
SELECT genres.name, COUNT(*) AS totalPeliculasXGenero FROM genres LEFT JOIN movies ON movies.genre_id = genres.id
GROUP BY genres.name HAVING totalPeliculasXGenero >= 3;


#5 el nombre y apellido actores todas las películas de la guerra de las galaxias, no se repitan.
SELECT DISTINCT actor_movie.actor_id AS actorId, actors.first_name, actors.last_name
FROM actor_movie 
INNER JOIN movies ON actor_movie.movie_id = movies.id
INNER JOIN actors ON actors.id = actor_movie.actor_id
WHERE movies.title LIKE "%Guerra de las galaxias%";


