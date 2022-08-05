-- 2 Agregar una película a la tabla movies.
INSERT INTO movies (title,rating,awards,release_date,length,genre_id) VALUES ('Interestelar',8.6,8,'20141026',169,5);
SELECT * FROM movies;
-- 3 Agregar un género a la tabla genres.
INSERT INTO genres (created_at,name,ranking,active) VALUES ('20220804','Distopia',13,1);
SELECT * FROM genres;
-- 4 Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies
SET genre_id=13
WHERE title='Interestelar';
SELECT * FROM movies;
-- 5 Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors
SET favorite_movie_id =22
WHERE first_name='Leonardo' AND last_name='Di Caprio';
SELECT * FROM actors;
-- 6 Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE pelisTemp AS (SELECT * FROM movies);
-- 7 Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM pelisTemp
WHERE awards<5;
-- 8 Obtener la lista de todos los géneros que tengan al menos una película.
SELECT generos.name as genero, count(*) as cantidad
FROM genres generos
	INNER JOIN movies peliculas
    ON generos.id=peliculas.genre_id
    GROUP BY generos.name HAVING cantidad>=1;
-- 9 Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT actores.first_name as nombre, actores.last_name as apellido, pelis.title as pelicula_favorita, pelis.awards as premios
FROM actors actores
	INNER JOIN movies pelis 
    ON pelis.id=actores.favorite_movie_id
    GROUP BY actores.id HAVING pelis.awards>3;
-- 10 Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

-- 11 ¿Qué son los índices? ¿Para qué sirven?

-- 12 Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX pelis_titulo_index ON movies(title);
-- 13 Chequee que el índice fue creado correctamente
SHOW INDEX FROM movies;