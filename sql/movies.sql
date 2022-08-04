Explicar el concepto de normalización y para que se utiliza.
Dicho brevemente, sería ordenar nuestra BD para que no haya redundancia de datos

Agregar una película a la tabla movies.
INSERT INTO movies(title, rating, awards, release_date, genre_id) 
VALUES('nahuel', 5, 10, (CAST('2015-12-25 15:32:06.427' AS DateTime)), 2)


Agregar un género a la tabla genres.
INSERT INTO genres(name, ranking, active)
VALUES(terror, 10, 1)

Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies as m
SET genre_id = 15
WHERE id = 22;

Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors as a 
SET favorite_movie_id = 22
WHERE id =3;

Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE prueba SELECT * FROM movies;

Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM prueba
WHERE awards < 5;

Obtener la lista de todos los géneros que tengan al menos una película.
SELECT g.name, m.title
FROM genres g
	JOIN movies m ON g.id = m.genre_id;

Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT a.first_name 
FROM actors a
	JOIN movies m ON m.id = a.favorite_movie_id
WHERE m.awards > 3;
Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

¿Qué son los índices? ¿Para qué sirven?
Sirven para hacer que SQL pueda acceder más rapidamente a recursos que suelen ser utilizados.
Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX movies_title_idx ON movies (title)
Chequee que el índice fue creado correctamente.
