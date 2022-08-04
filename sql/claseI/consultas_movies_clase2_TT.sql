#CLASE 2 PARTE2
#1.Explicar el concepto de normalización y para que se utiliza.
#el proceso de normalizacion lleva acabo la organizacion y optimizacion de los datos aplicando algunas reglas en su implementacion
#2.Agregar una película a la tabla movies.
USE movies_db;
INSERT INTO movies (title, rating, awards, release_date, length, genre_id)
VALUES('Sueno de fuga', 9.5, 2, '1994-09-22 00:00:00', 142,3);
#3.Agregar un género a la tabla genres.
INSERT INTO genres (name, ranking, active)
VALUES ('Juvenil', 13, 1);
#4.Asociar a la película del Ej 2. con el género creado en el Ej. 3.
-- SET SQL_SAFE_UPDATE = 0;
UPDATE movies
SET genre_id = 13
WHERE id = 22;
#5.Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors
SET favorite_movie_id = 22
WHERE id = 12;
#6.Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_tmp
SELECT * FROM movies;
#7.Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
-- SET SQL_SAFE_UPDATES = 0;
DELETE FROM movies_tmp
WHERE awards < 5;
SELECT * FROM movies_tmp;
#8.Obtener la lista de todos los géneros que tengan al menos una película.
SELECT DISTINCT g.name AS generos
FROM movies AS m
INNER JOIN genres AS g ON m.genre_id = g.id;
#9.Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT a.first_name AS nombre, a.last_name AS apellido
FROM actors AS a
INNER JOIN movies AS m ON a.favorite_movie_id = m.id
WHERE m.awards > 3;
#10.Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
EXPLAIN ANALYZE SELECT DISTINCT g.name AS generos
FROM movies AS m
INNER JOIN genres AS g ON m.genre_id = g.id;
EXPLAIN ANALYZE SELECT a.first_name AS nombre, a.last_name AS apellido
FROM actors AS a
INNER JOIN movies AS m ON a.favorite_movie_id = m.id
WHERE m.awards > 3;
#11.¿Qué son los índices? ¿Para qué sirven?
#el indice es una estructura de datos que mejora el rendimiento de las consultas por medio de un identificador unico para cada fila de una tabla
#12.Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX movies_title_idx
	ON movies(title);
#13.Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;