-- SQL 1 - MOVIES DB

-- 1) Abrir desde MySql Workbench el archivo movies_db.sql y ejecutar su contenido. 
-- OK

-- 2) Mostrar todos los registros de la tabla de movies.
SELECT * FROM movies;

-- 3) Mostrar el nombre, apellido y rating de todos los actores.
SELECT first_name, last_name, rating 
FROM actors;

-- 4) Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español
SELECT title as Título 
FROM series;

-- 5) Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
SELECT first_name, last_name
FROM actors
WHERE rating > 7.5; 

-- 6) Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor 
-- a 7.5 y con más de dos premios.
SELECT title, rating, awards
FROM movies
WHERE rating > 7.5 AND awards > 2;

-- 7) Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
SELECT title, rating
FROM movies
ORDER BY rating;

-- 8) Mostrar los títulos de las primeras tres películas en la base de datos.
SELECT title FROM movies LIMIT 3;

-- 9) Mostrar el top 5 de las películas con mayor rating.
SELECT title, rating 
FROM movies
ORDER BY rating DESC
LIMIT 5;

-- 10) Mostrar las top 5 a 10 de las películas con mayor rating.
SELECT title, rating 
FROM movies
ORDER BY rating DESC
LIMIT 5 OFFSET 5;

-- 11) Listar los primeros 10 actores (sería la página 1)
SELECT id, CONCAT(first_name, " ", last_name) as actor
FROM actors
LIMIT 10;

-- 12) Luego usar offset para traer la página 3
SELECT id, CONCAT(first_name, " ", last_name) as actor
FROM actors
LIMIT 10 OFFSET 20;

-- 13) Hacer lo mismo para la página 5
SELECT id, CONCAT(first_name, " ", last_name) as actor
FROM actors
LIMIT 10 OFFSET 40;

-- 14) Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
SELECT title, rating
FROM movies
WHERE title LIKE "Toy Story%";

-- 15) Mostrar a todos los actores cuyos nombres empiezan con Sam.
SELECT CONCAT(first_name, " ", last_name) as actor
FROM actors
WHERE first_name LIKE "Sam%";

-- 16) Mostrar el título de las películas que salieron entre el 2004 y 2008.
SELECT title, YEAR(release_date) as release_year
FROM movies
WHERE YEAR(release_date) BETWEEN 2004 AND 2008;

-- 17) Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.
SELECT title, rating, awards, YEAR(release_date) as release_year
FROM movies
WHERE rating > 3
	  AND awards > 1 
      AND YEAR(release_date) BETWEEN 1988 AND 2009
ORDER BY rating;

-- 18) Traer el top 3 a partir del registro 10 de la consulta anterior.
SELECT title, rating, awards, YEAR(release_date) as release_year
FROM movies
WHERE rating > 3
	  AND awards > 1 
      AND YEAR(release_date) BETWEEN 1988 AND 2009
ORDER BY rating
LIMIT 3 OFFSET 10;