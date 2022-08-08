USE movies_db;

-- PUNTO 2
INSERT INTO movies (title, rating, awards, release_date, length, genre_id)
VALUES ('Como entrenar a tu dragón', 9.5, 3, '2010-03-26 00:00:00', 100, NULL);

-- PUNTO 3
INSERT INTO genres (created_at, name, ranking, active)
VALUES ('2010-03-26 00:00:00', 'Independiente', 13, 1);

-- PUNTO 4
UPDATE movies
SET genre_id = 13
WHERE id = 22;

-- PUNTO 5
UPDATE actors
SET favorite_movie_id = 22
WHERE id = 3;

-- PUNTO 6
CREATE TEMPORARY TABLE movies_tmp AS (SELECT * FROM movies);

-- PUNTO 7
DELETE FROM movies_tmp
WHERE awards < 5;

-- PUNTO 8
SELECT g.name as nombre, COUNT(*) as total_peliculas FROM movies m
INNER JOIN genres g
ON m.genre_id = g.id
GROUP BY g.name
HAVING total_peliculas >= 1;

-- PUNTO 9
SELECT CONCAT(a.first_name, ' ', a.last_name) as nombre_actor, title as pelicula_favorita, awards as premios
FROM actors a
INNER JOIN movies m
ON a.favorite_movie_id = m.id
WHERE m.awards > 3;

-- PUNTO 10
EXPLAIN SELECT * FROM movies;
EXPLAIN DELETE FROM movies_tmp WHERE awards < 5;

-- PUNTO 12
CREATE INDEX title_idx ON movies(title);

-- PUNTO 13
SHOW INDEX FROM movies;

-- CONSULTAS
SELECT * FROM genres;
SELECT * FROM movies;
SELECT * FROM actors;
SELECT * FROM movies_tmp;