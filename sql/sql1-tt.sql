-- Todos los registros de la tabla movie
SELECT *
FROM movies_db.movies;

-- Mostrar nombre, apellido y rating de los actores
SELECT first_name, last_name, rating
FROM movies_db.actors;

-- Mostrar titulo d serie con alias
SELECT title titulo
FROM movies_db.series;

-- Mostrar nombre y apellido de los actores con rating mayor a 7.5
SELECT first_name, last_name, rating
FROM movies_db.actors
WHERE rating > 7.5;

-- Mostrar titulo de peliculas, rating y premios con rating mayor a 7.5 y mas de 2 premios
SELECT title, rating, awards
FROM movies_db.movies
WHERE rating > 7.5 AND awards > 2;

-- Mostrar el titulo de las peliculas y el rating ordenado ASC
SELECT title, rating
FROM movies_db.movies
ORDER BY rating ASC;

-- Mostrar el titulo de las primeras 3 peliculas en la BD
SELECT title
FROM movies_db.movies
LIMIT 3;

-- Mostrar top 5 de las peliculas con mayor rating
SELECT title, rating
FROM movies_db.movies
ORDER BY rating DESC
LIMIT 5;

-- Mostrar top 5 a 10 de las peliculas con mayor rating
SELECT title, rating
FROM movies_db.movies
ORDER BY rating DESC
LIMIT 5 OFFSET 10;

-- Listar los primeros 10 actores (Seria la pagina 1)
SELECT *
FROM movies_db.actors
LIMIT 10;

-- Luego usar offset para traer la página 3
SELECT *
FROM movies_db.actors
LIMIT 10 OFFSET 20;

-- Hacer lo mismo para la página 5
SELECT *
FROM movies_db.actors
LIMIT 10 OFFSET 40;

-- Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
SELECT title, rating
FROM movies_db.movies
WHERE title LIKE '%Toy Story%';

-- Mostrar a todos los actores cuyos nombres empiezan con Sam.
SELECT *
FROM movies_db.actors
WHERE first_name LIKE 'Sam%';

 -- Mostrar el título de las películas que salieron entre el 2004 y 2008.
SELECT title, release_date, extract(year from release_date)
FROM movies_db.movies
WHERE extract(year from release_date) BETWEEN '2004' AND '2008';

-- Traer el título de las películas con el rating mayor a 3, 
-- con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. 
-- Ordenar los resultados por rating.
SELECT title, rating, awards, extract(year from release_date)
FROM movies_db.movies
WHERE rating > 3 AND awards > 1 AND extract(year from release_date) BETWEEN '1988' AND '2009'
ORDER BY rating DESC;

 -- Traer el top 3 a partir del registro 10 de la consulta anterior.
SELECT title, rating, awards, extract(year from release_date)
FROM movies_db.movies
WHERE rating > 3 AND awards > 1 AND extract(year from release_date) BETWEEN '1988' AND '2009'
ORDER BY rating DESC
LIMIT 3 OFFSET 10;