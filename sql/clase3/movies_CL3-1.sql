USE movies_db;
#Explicar el concepto de normalización y para que se utiliza.
# Estandarización y validación de datos que consiste en eliminar redundancias e inconsistencias de datos
# con reglas que protegen la integridad de la info así como su gestión.

#Agregar una película a la tabla movies.
INSERT INTO movies(title, rating, awards, length, release_date) VALUES("Orgullo y Prejuicio",7.8,2,127, "2006-02-10");
#Agregar un género a la tabla genres.
INSERT INTO genres(name,active, ranking) VALUES("Romance", 1,13);
#Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies SET genre_id = 14 WHERE id = 22;
#Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors SET favorite_movie_id = 22 WHERE id = 32;
#Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_temp AS(SELECT * FROM movies);
SELECT * FROM movies_temp;
#Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES=0;
DELETE FROM movies_temp WHERE awards < 5;
#Obtener la lista de todos los géneros que tengan al menos una película.
SELECT name FROM genres WHERE id IN (SELECT genre_id FROM movies) GROUP BY name;
#Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT first_name, last_name FROM actors WHERE favorite_movie_id IN (SELECT id FROM movies WHERE awards>3);
#Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
EXPLAIN DELETE FROM movies_temp WHERE awards < 5;
# Solo pude hacer el EXPLAIN para la query de DELETE, para la de CREATE TEMPORARY no.alter
#¿Qué son los índices? ¿Para qué sirven?
# Son un mecanismos para optimizar consultas SQL, mejoran el acceso a los datos
#Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX title_idx ON movies(title);
#Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;
