USE movies_db;

# EJERCICIO 2
INSERT INTO movies(created_at, updated_at, title, rating, awards, release_date, length, genre_id)
VALUES('2022-04-21 00:00:00','2022-08-11 00:00:00','Ted', 8.5, 4, '2022-03-21', 150, 1);

# EJERCICIO 3
INSERT INTO genres(created_at, updated_at, name, ranking, active)
VALUES('2022-04-21 00:00:00','2022-08-11 00:00:00','Romance', 13, 1);

# EJERCICIO 4
UPDATE movies
SET genre_id = 13
WHERE id = 22;

# EJERCICIO 5
UPDATE actors
SET favorite_movie_id = 22
WHERE id = 1;

# EJERCICIO 6
CREATE TEMPORARY TABLE movies_tmp AS (SELECT * FROM movies);

# EJERCICIO 7
DELETE FROM movies_tmp 
WHERE awards < 5;

# EJERCICIO 8
SELECT g.name as genero,  COUNT(*) as total_peliculas FROM movies m
INNER JOIN genres g
ON m.genre_id = g.id
GROUP BY g.name
HAVING total_peliculas >= 1;

# EJERCICIO 9
SELECT CONCAT(a.first_name, ' ', a.last_name) as nombre_completo, m.title as pelicula, m.awards as premios FROM actors a
INNER JOIN movies m
ON m.id = a.favorite_movie_id
WHERE m.awards > 3;

# EJERCICIO 10
EXPLAIN SELECT * FROM movies;
EXPLAIN DELETE FROM movies_tmp WHERE awards < 5;

# EJERCICIO 12
CREATE INDEX title_idx ON movies(title);

# EJERCICIO 13
SHOW INDEX FROM movies;