-- 1) Explicar el concepto de normalización y para que se utiliza.
-- R: Conjunto de etapas, llamadas formas normales, que se van aplicando en el diseño de una BD 
-- para evitar la redundancia y proteger la integridad de los datos, además de facilitar el acceso y reducir 
-- el tiempo de ejecución de consultas (se puede aplicar hasta la 6FN pero lo usual es hasta la 3FN).
-- 1FN --> No existen filas repetidas y todos los atributos son atómicos.
-- 2FN --> Todos los atributos que no forman parte de la clave principal tienen dependencia funcional completa de ella.
-- 3FN --> No existen dependencias transitivas (cuando una columna depende de otra que no es clave principal)

-- 2) Agregar una película a la tabla movies.
INSERT INTO movies (id, created_at, updated_at, title, rating, awards, release_date, length)
VALUES (22, NULL, NULL, 'El Discurso del Rey', 8.0, 4, '2010-11-26 00:00:00', 120);

-- 3) Agregar un género a la tabla genres.
INSERT INTO genres (id, created_at, updated_at, name, ranking, active) 
VALUES (13,'2022-08-04 23:00:00', NULL, 'Historico', 13, 1);

-- 4) Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies SET genre_id = 13 WHERE id = 22;

SELECT movies.title AS Película, genres.name AS Género
FROM movies JOIN genres ON movies.genre_id = genres.id
WHERE movies.title = 'El Discurso del Rey';

-- 5) Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors SET favorite_movie_id = 22 WHERE id = 5;

SELECT actors.id, CONCAT(actors.first_name, " ", actors.last_name) AS Actor, movies.title AS Película
FROM movies JOIN actors ON actors.favorite_movie_id = movies.id
WHERE actors.id = 5;

-- 6) Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_copy
			 SELECT * FROM movies;

SELECT * FROM movies_copy;

-- 7) Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
DELETE FROM movies_copy WHERE awards < 5;

SELECT * FROM movies_copy;

-- 8) Obtener la lista de todos los géneros que tengan al menos una película.
SELECT DISTINCT genres.name AS Género 
FROM genres JOIN movies ON genres.id = movies.genre_id;

-- (Otra forma)
SELECT genres.name AS Género, COUNT(movies.title) AS Películas
FROM genres JOIN movies ON genres.id = movies.genre_id
GROUP BY genres.name;

-- 9) Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT CONCAT(actors.first_name, " ", actors.last_name) AS Actor, 
movies.title AS "Película Favorita", movies.awards AS Premios
FROM movies JOIN actors ON actors.favorite_movie_id = movies.id
WHERE movies.awards > 3;

-- 10) Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
-- R: Resultado en capturas adjuntas.

-- 11) ¿Qué son los índices? ¿Para qué sirven?
-- R: Los índices son una estructura de datos que permiten un acceso más rápido a los registros de una tabla
-- en una base de datos. Sirve para mejorar el rendimiento de consultas que son frecuentes.
-- No obstante, no son recomendados en aquellas tablas en las que se utilizan operaciones de escritura 
-- frecuentes (INSERT, DELETE, UPDATE), esto es porque los índices se actualizan cada vez que se
-- modifica una columna.

-- 12) Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX title_idx ON movies(title);

-- 13) Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;