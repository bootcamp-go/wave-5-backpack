-- 2 
INSERT INTO movies (created_at, updated_at, title, rating, awards, release_date, length, genre_id)
VALUES (null, null, "Parque Jurasico II", 7.0, 1, "2003-04-04 00:00:00", 270, 5);

-- 3
INSERT INTO genres (created_at, updated_at, name, ranking, active)
VALUES ("2019-07-04 00:00:00", null, "Otros", 13, 1);

-- 4
UPDATE movies
SET genre_id = 13
WHERE id = 22;

-- 5
UPDATE actors
SET favorite_movie_id = 22
WHERE id = 49;

-- 6
-- DROP TABLE tempTable
CREATE TEMPORARY TABLE tempTable
SELECT * FROM movies;

-- 7
-- SET SQL_SAFE_UPDATES = 1;
DELETE FROM tempTable 
WHERE awards < 5;

-- 8
SELECT g.name Genero, COUNT(*) CantPeliculas
FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.name 
HAVING CantPeliculas > 1;

-- 9
SELECT a.* -- first_name Nombre, a.last_name Apellido, m.title PeliFavorita
FROM actors a
INNER JOIN actor_movie am ON am.actor_id = a.id
INNER JOIN movies m ON am.movie_id = m.id
WHERE m.awards > 3;

-- 10
EXPLAIN SELECT * FROM tempTable;
EXPLAIN DELETE FROM tempTable WHERE awards < 5;

-- 12
CREATE INDEX title_idx 
ON movies (title);

-- 13
SHOW INDEX FROM movies