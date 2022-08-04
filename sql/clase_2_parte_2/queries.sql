USE movies_db;

# Mostrar el título y el nombre del género de todas las series.
SELECT s.title, g.name FROM series s 
INNER JOIN genres g ON s.genre_id = g.id;

# Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT a.first_name, a.last_name, m.title FROM actor_movie am 
INNER JOIN actors a ON a.id = am.actor_id
INNER JOIN movies m ON m.id = am.movie_id;

# Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT s.title,COUNT(*) as total_seasons FROM seasons se
INNER JOIN series s ON s.id = se.serie_id
GROUP BY s.title;

# Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT g.name, COUNT(*) as total_movies FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.name
HAVING total_movies >= 3;

# Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT a.first_name, a.last_name FROM actor_movie am 
INNER JOIN actors a ON a.id = am.actor_id
INNER JOIN movies m ON m.id = am.movie_id
WHERE m.title like "%Guerra de las galaxias%"
GROUP BY a.first_name, a.last_name
HAVING COUNT(*) = (SELECT COUNT(*) FROM movies m
WHERE m.title like "%Guerra de las galaxias%");