-- Explicar el concepto de normalización y para que se utiliza.
-- Es un proceso para eliminar la redundancia  e inconsistencia de datos , protegiendo su integridad y favoreciendo 
-- la interpretacion para que sea mas simple de consultar y su gestión sea más eficiente.


-- Agregar una película a la tabla movies.
INSERT INTO `movies_db`.`movies`
(`title`,
`rating`,
`awards`,
`release_date`,
`length`,
`genre_id`)
VALUES('Fast and Furious',9,3,'2012-05-04 00:00:00',90,1);

-- Agregar un género a la tabla genres.
INSERT INTO `movies_db`.`genres`
(`created_at`,
`name`,
`ranking`,
`active`)
VALUES
('2011-07-04 00:00:00','Street Racing',13,1);
-- Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies SET genre_id=13 WHERE id=22;
-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors SET favorite_movie_id=22 WHERE ID=2;
-- Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_temp as (SELECT * FROM movies);
-- Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
DELETE FROM movies_temp WHERE awards<5;
SET SQL_SAFE_UPDATES = 1;
-- Obtener la lista de todos los géneros que tengan al menos una película.
SELECT * FROM genres where id in (SELECT genre_id FROM MOVIES);
-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards
SELECT * FROM actors WHERE favorite_movie_id IN (SELECT id FROM movies WHERE awards>3);
-- Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

	EXPLAIN SELECT * FROM movies;
	EXPLAIN DELETE FROM movies_temp WHERE awards<5;

-- ¿Qué son los índices? ¿Para qué sirven?
-- Se utiliza este mecanismo para optimizar las búsquedas por SQL. 
-- Mejoran los tiempos de respuesta y acceso a datos utilizando
-- una ruta mas directa a los mismos.
-- Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX movies_title_idx ON movies(title);
-- Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;


