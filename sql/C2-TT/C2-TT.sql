USE movies_db;

INSERT INTO movies
(title, rating, awards,release_date,length,genre_id)
VALUES
('Rapido y Furioso: reto Tokio',8.5,1,'2006-10-04',104,4);

INSERT INTO genres
(name, ranking, active)
VALUES
('Autos',13,1);

UPDATE movies
SET genre_id = 13
WHERE id = 22;

UPDATE actors
SET favorite_movie_id = 13
WHERE id IN (4,21);

CREATE TEMPORARY TABLE movies_temp 
SELECT * FROM movies;

DELETE FROM movies_temp WHERE awards < 5;

SELECT gn.name, COUNT(mo.genre_id) as total_movies
FROM genres gn
INNER JOIN movies mo
ON gn.id = mo.genre_id
GROUP BY gn.name HAVING COUNT(mo.genre_id)>1;

SELECT ac.first_name, ac.last_name, mo.title, mo.awards
FROM actors ac
INNER JOIN movies mo
ON ac.favorite_movie_id = mo.id
GROUP BY ac.first_name, ac.last_name,mo.title, mo.awards HAVING mo.awards > 3; 

CREATE INDEX movies_title_idx
ON movies (title);

SHOW INDEX FROM movies;