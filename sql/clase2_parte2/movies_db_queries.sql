use movies_db;
# 2 Exercise
# Agregar una película a la tabla movies.
INSERT INTO movies 
	(title, rating, awards, release_date, length, genre_id) 
VALUES
	("Monster Inc", 9.0, 5, "2001-12-21", 120, 10);

# 3 Exercise
# Agregar un género a la tabla genres.
INSERT INTO genres 
	(created_at, name, ranking, active) 
VALUES
	("2013-07-03", "Historicas", 13, 1);

# 4 Exercise
# Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies
SET
	title = "Togo",
	genre_id = 14
WHERE
	id = 22;
    
# 5 Exercise
# Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors
SET
	favorite_movie_id = 22
WHERE 
	id = 49;

# 6 Exercise
# Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_temporal AS (SELECT * FROM movies);

# 7 Exercise
# Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

-- Execute bdefore to disable safe update mode
SET SQL_SAFE_UPDATES=0;

DELETE 
FROM movies_temporal
WHERE
	awards < 5;
    
# 8 Exercise
# Obtener la lista de todos los géneros que tengan al menos una película.
SELECT genres.* 
FROM genres
INNER JOIN movies ON movies.genre_id = genres.id
GROUP BY genres.id;

# 9 Exercise
# Obtener la lista de actores cuya película favorita haya ganado más de 3 awards
SELECT actors.*, movies.title, movies.awards
FROM actors 
INNER JOIN movies ON movies.id = actors.favorite_movie_id
WHERE movies.awards > 3;

# 10 Exercise
# Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
DROP TABLE movies_temporal;
CREATE TABLE movies_temporal AS (SELECT * FROM movies);

# 12 Exercise
# Crear un índice sobre el nombre en la tabla movies.
ALTER TABLE movies 
ADD INDEX (title);

# 13 Exercise
# Crear un índice sobre el nombre en la tabla movies.
SHOW INDEX FROM movies;


# 2, 3, 4, 5
SELECT * FROM movies WHERE awards > 3;
SELECT count(*) FROM movies_temporal;
SELECT * FROM actors WHERE ;
SELECT count(*) FROM genres;
