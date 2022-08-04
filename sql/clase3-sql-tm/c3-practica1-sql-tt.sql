/*-----------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  SQL 2 - MOVIES DB
         Author:	Israel Fabela
	   Language:	mysql  Ver 8.0.29 for macos12.2 on arm64
		  Topic:	Base de Datos - SQL

	© Mercado Libre - IT Bootcamp 2022

-------------------------------------------------------------------------*/

-- Utilizando la tabla 'movies_db'
USE movies_db;

-- 2. Agregar una película a la tabla movies.
INSERT INTO `movies` (title,rating,awards,release_date,length,genre_id)
	VALUES ('Encanto', 9.1, 12, '2021-11-24 00:00:00', 109, 10);

-- 3. Agregar un género a la tabla genres.
INSERT INTO `genres` (created_at, name, ranking, active)
	VALUES ('2005-07-03 22:00:00', 'Terror', 13, 1);

-- 4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies_db.movies
	SET genre_id = 13 
	WHERE title = "Encanto" AND id <> 0;
    
-- 5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors
	SET favorite_movie_id = 13
    WHERE first_name = 'Johnny' AND  last_name = 'Depp' AND id <> 0;

-- 6. Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE tmp_movies AS (SELECT * FROM movies);

-- 7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
DELETE FROM tmp_movies WHERE awards < 5;
SELECT * FROM tmp_movies;
SET SQL_SAFE_UPDATES = 1;

-- 8. Obtener la lista de todos los géneros que tengan al menos una película.
SELECT * FROM genres WHERE id IN (SELECT genre_id FROM MOVIES);

-- 9. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT * FROM actors WHERE favorite_movie_id IN (SELECT id FROM movies WHERE awards > 3);

-- 12. Crear un índice sobre el nombre en la tabla movies
CREATE INDEX title_movies_index ON movies(title);

-- 13. Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;

