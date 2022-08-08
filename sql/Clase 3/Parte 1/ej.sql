-- 2.
INSERT INTO movies 
(id, title, rating, awards, release_date, length, genre_id) 
VALUES (22, 'Deadpool', 9.2, 3, STR_TO_DATE('2018-04-12', '%Y-%m-%d'), 100, 4);
-- 3 
INSERT INTO genres
(id, name, ranking)
VALUES (13, 'Violencia', 13);
-- 4
UPDATE movies m
SET m.genre_id = 13
WHERE m.id = 22;
-- 5
UPDATE actors a
SET a.favorite_movie_id = 22
WHERE a.id = 3;
-- 6 
CREATE TEMPORARY TABLE temp_movies (SELECT * FROM movies);
-- 7
DELETE FROM temp_movies tmpm WHERE tmpm.awards < 5;
-- 8 
SELECT g.id, g.name, COUNT(m.id) conteo
FROM genres g
JOIN movies m
WHERE g.id = m.genre_id
GROUP BY m.genre_id
HAVING conteo >= 1;
-- 9
SELECT a.first_name, a.last_name, m.awards
FROM actors a
JOIN movies m
WHERE a.favorite_movie_id = m.id
AND m.awards > 3;
-- 12
ALTER TABLE movies
ADD INDEX IDX_movies_title (title);
-- 13
SHOW INDEX FROM movies;
