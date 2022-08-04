/*
1- Explicar el concepto de normalización y para que se utiliza.
	Rta: La normalización es un proceso de estandarización y validación de datos que consiste en eliminar las redundancias o inconsistencias, completando datos mediante una serie de reglas que actualizan la información, protegiendo su integridad y favoreciendo la interpretación, para que así sea más simple de consultar y más eficiente para quien la gestiona.
*/
# 2- Agregar una película a la tabla movies.

INSERT INTO movies (id, title, rating, awards, release_date, genre_id) 
VALUES (22, "El lobo de Wall Street", 99, 4, "2014-11-04 00:00:00", 3);

# 3- Agregar un género a la tabla genres.

INSERT INTO genres (id, name, ranking)
VALUES (13, "Thriller", 13);

# 4- Asociar a la película del Ej 2. con el género creado en el Ej. 3.

UPDATE movies SET genre_id = 13
WHERE id = 22;

# 5- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors SET favorite_movie_id = 22
WHERE id = 4;

# 6- Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_temp
AS (SELECT * FROM movies);

# 7- Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM movies_temp
WHERE awards < 5;

# 8- Obtener la lista de todos los géneros que tengan al menos una película.
SELECT g.name FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.name;
  
# 9- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT a.first_name, a.last_name, m.awards FROM actors a
INNER JOIN movies m ON m.id = a.favorite_movie_id
WHERE m.awards > 3;

# 10- Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

# 11- ¿Qué son los índices? ¿Para qué sirven?
/*
- Son un mecanismo para optimizar consultas en SQL.
- Mejoran sustancialmente los tiempos de respuesta en Queries complejas.
- Mejoran el acceso a los datos al proporcionar una ruta más directa a los registros. 
- Evitan realizar escaneos (barridas) completas o lineales de los datos en una tabla.
*/

# 12- Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX movies_title_idx
ON movies (title);

# 13- Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;