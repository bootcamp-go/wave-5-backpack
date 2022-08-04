USE movies_db;

# Explicar el concepto de normalización y para que se utiliza.
# Se utilizar para eliminar duplicidad, inconsistencias y redundancias. se aplica las 3 Formas normales para lograr esaa normalizacion.

# Agregar una película a la tabla movies.
INSERT INTO movies(title, rating, awards, release_date, length, genre_id)
VALUES('THOR THUNDERBOLT', 7, 3, '2022-06-25', 150, 4);

# Agregar un género a la tabla genres.
INSERT INTO genres(created_at, name, ranking, active)
VALUES (NOW(), 'Policial', 13, true);

# Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies m
SET genre_id = LAST_INSERT_ID()
where m.id = 22;

# Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors
SET favorite_movie_id = 22
where id = 1;

# Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE temp_movies 
SELECT * FROM movies;

# Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM temp_movies
WHERE temp_movies.awards < 5;

# Obtener la lista de todos los géneros que tengan al menos una película.
SELECT g.name FROM genres g
INNER JOIN movies m ON g.id = m.genre_id
GROUP BY g.name
HAVING COUNT(*) > 0;
# Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT a.first_name, a.last_name FROM actors a
INNER JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3
GROUP BY a.first_name, a.last_name
HAVING COUNT(*) > 0;

# Utilizar el explain plan para analizar las consultas
EXPLAIN SELECT a.first_name, a.last_name FROM actors a
INNER JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3
GROUP BY a.first_name, a.last_name
HAVING COUNT(*) > 0;

EXPLAIN SELECT a.first_name, a.last_name FROM actors a
INNER JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3
GROUP BY a.first_name, a.last_name
HAVING COUNT(*) > 0;

# ¿Qué son los índices? ¿Para qué sirven? , los indices son referencias a columnas en tablas para hacer busquedas de forma más eficiente de un campo sin indexar.

# Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX title_idx
ON movies(title);

# Chequee que el índice fue creado correctamente.
SHOW KEYS FROM movies;
