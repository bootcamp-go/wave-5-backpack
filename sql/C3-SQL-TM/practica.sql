-- 1. Explicar el concepto de normalización y para que se utiliza.

--    Es un proceso de estandarización y validación de datos, que nos permite reducir o evitar la redundancia de datos

-- 2. Agregar una película a la tabla movies.

/*
      INSERT INTO movies (title,rating,awards,release_date,length) VALUES ('A prueba de fuego',8.5,2,'2008-09-26',122);
      SELECT * FROM movies WHERE id = 22;
*/

-- 3. Agregar un género a la tabla genres.

/*
      INSERT INTO genres (created_at,name,ranking,active) VALUES ('2015-07-03 22:00:00','Religiosas',13,1);
      SELECT * FROM genres WHERE id = 13;
*/

-- 4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.

/*
      UPDATE movies SET genre_id = 13 WHERE id = 22;
      SELECT * FROM movies WHERE id = 22;
*/

-- 5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.

/*
      UPDATE actors SET favorite_movie_id = 22 WHERE id = 47;
      SELECT * FROM actors WHERE id = 47;
*/

-- 6. Crear una tabla temporal copia de la tabla movies.

/*
      CREATE TEMPORARY TABLE movies_temp AS (SELECT * FROM movies);
      SELECT * FROM movies_temp;
*/

-- 7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

/*
      DELETE FROM movies_temp WHERE awards < 5;
      SELECT * FROM movies_temp;
*/

-- 8. Obtener la lista de todos los géneros que tengan al menos una película.

/*
      SELECT g.name, COUNT(g.id) AS total_peliculas
      FROM genres g
      LEFT JOIN movies m ON g.id = m.genre_id
      GROUP BY g.name
*/

-- 9. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 

/*
      SELECT a.first_name AS nombre, a.last_name AS apellido, m.title, m.awards
      FROM actors a
      INNER JOIN movies m ON a.favorite_movie_id = m.id
      WHERE m.awards > 3
      ORDER BY m.awards
*/

-- 10. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

/*
       EXPLAIN DELETE FROM movies_temp WHERE awards < 5;
*/

-- 11. ¿Qué son los índices? ¿Para qué sirven?

--     Es un mecanismo para optimizar las consultas SQL a través de la generación de rutas directas a la información,
--     para evitar escaneos completos de los datos

-- 12. Crear un índice sobre el nombre en la tabla movies.

/*
       CREATE INDEX title_idx ON movies(title);
*/

-- 13. Chequee que el índice fue creado correctamente.

/*
       SHOW INDEX FROM movies;
*/
