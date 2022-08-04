USE  movies_db;

# -------- 1 Exercise --------
# Mostrar el título y el nombre del género de todas las series.
SELECT s.title as title_serie, g.name as name_genre FROM series s INNER JOIN genres g ON g.id = s.id;

# -------- 2 Exercise --------
# Mostrar el título y el nombre del género de todas las series.
SELECT episodes.title, actors.first_name, actors.last_name 
FROM episodes 
INNER JOIN actor_episode ON actor_episode.episode_id = episodes.id 
INNER JOIN actors ON actors.id = actor_episode.episode_id;

# -------- 3 Exercise --------
# Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT series.title, SUM(seasons.id) 
FROM seasons 
INNER JOIN series ON series.id = seasons.id
GrOUP BY series.title;

# -------- 4 Exercise --------
# Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT genres.name, SUM(movies.id) as cantity_movies_by_genre 
FROM movies
INNER JOIN genres ON genres.id = movies.genre_id
GROUP BY genres.name 
HAVING cantity_movies_by_genre >= 3;

# -------- 5 Exercise --------
# Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT actors.first_name, actors.last_name
FROM actors 
INNER JOIN actor_movie ON actor_movie.actor_id = actors.id 
INNER JOIN movies ON movies.id = actor_movie.movie_id
WHERE title LIKE "%la guerra de las galaxias%";