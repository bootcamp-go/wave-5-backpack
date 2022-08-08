USE movies_db;
-- 1 Explicar el concepto de normalización y para que se utiliza.
-- 	Es un proceso de estandarización y validación de datos que consisten en eliminar redundancias o incosistencias. Permite mejorar la integridad, interpretación y la eficiencia de las consultas

-- 2 Agregar una película a la tabla movies.
	INSERT INTO movies (id, created_at, updated_at, title, rating, awards, release_date, length, genre_id)
    VALUES (DEFAULT, DEFAULT, DEFAULT, 'nueva peli', 7.8, 3, '1999-01-04 00:00:00', 190, 2);
    
-- 3 Agregar un género a la tabla genres.
	INSERT INTO genres (id, created_at, updated_at, name, ranking, active)
    VALUES (DEFAULT, DEFAULT, DEFAULT, 'nuevo_genero', 13, DEFAULT);
    
-- 4 Asociar a la película del Ej 2. con el género creado en el Ej. 3.
	UPDATE movies
    SET genre_id = 13
    WHERE id = 22 AND title = 'nueva peli'; 
    
-- 5 Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
	UPDATE actors
    SET favorite_movie_id = 22
    WHERE id = 4;
    
-- 6 Crear una tabla temporal copia de la tabla movies.
	CREATE TEMPORARY TABLE movies_tmp SELECT * FROM movies;

-- 7 Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
	SET SQL_SAFE_UPDATES = 0;
	DELETE FROM movies_tmp
    WHERE awards < 5;

-- 8 Obtener la lista de todos los géneros que tengan al menos una película.
	SELECT name as genero, count(*) as cant_peliculas FROM genres gr
    INNER JOIN movies mo ON gr.id = mo.genre_id
    GROUP BY name;

-- 9 Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
	SELECT first_name, last_name
    FROM actors ac
    INNER JOIN movies mo ON ac.favorite_movie_id = mo.id
    WHERE mo.awards > 3;

-- 10 Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
	-- se hace directamente apretando el rayo+lupa parado en la consulta

-- 11 ¿Qué son los índices? ¿Para qué sirven?
	-- Son un mecanismo para optimizar consultas en SQL.
	-- Mejoran sustancialmente los tiempos de respuesta en Queries complejas.
	-- Mejoran el acceso a los datos al proporcionar una ruta más directa a los registros.
	-- Evitan realizar escaneos (barridas) completas o lineales de los datos en una tabla.

-- 12 Crear un índice sobre el nombre en la tabla movies.
	CREATE INDEX movies_title_idx ON movies (title);

-- 13 Chequee que el índice fue creado correctamente.
	SHOW INDEX FROM movies;
