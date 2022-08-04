INSERT INTO movies_db.movies
	(
		created_at,
		updated_at,
		title,
		rating,
		awards,
		release_date,
		length,
		genre_id
		)
VALUES
	(
		null,
		null,
		'Casper',
		9.0,
		0,
		'1995-05-26 00:00:00',
		100,
		10
	);

INSERT INTO movies_db.genres
	(
		created_at,
		updated_at,
		name,
		ranking,
		activ
	)
VALUES
	(
		'2022-08-04 00:00:00',
		null,
		'En la vida real',
		13,
		1
	);

UPDATE movies_db.movies
SET	genre_id = 14
WHERE id = 22;

UPDATE movies_db.actors
SET	favorite_movie_id = 22
WHERE id = 1;

CREATE TEMPORARY TABLE movies_db.temp_movies 
	SELECT * FROM movies_db.movies;

DELETE FROM movies_db.temp_movies
WHERE awards < 5;

SELECT * FROM movies_db.temp_movies;

SELECT movies_db.g.name AS Nombre_Genero
FROM movies_db.genres AS g
JOIN movies_db.movies AS m
ON movies_db.g.Id = movies_db.m.genre_id
GROUP BY Nombre_Genero HAVING count(*) > 1;

SELECT movies_db.a.first_name AS Nombre_Actor,
	   movies_db.a.last_name AS Apellido_Actor,
       m.awards
FROM movies_db.actors AS a
JOIN movies_db.movies AS m
ON a.favorite_movie_id = m.Id
WHERE m.awards > 3;

CREATE INDEX movie_title_index 
ON movies_db.movies(title);

SHOW INDEX FROM movies_db.movies










