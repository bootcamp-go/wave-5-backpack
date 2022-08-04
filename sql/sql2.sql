Mostrar todos los registros de la tabla de movies.

SELECT * FROM movies;

Mostrar el nombre, apellido y rating de todos los actores.

SELECT first_name as nombre, last_name as apellido, rating FROM actors;

Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español

SELECT title as titulo FROM series

Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.

SELECT first_name as nombre, last_name as apellido FROM actors WHERE rating < 7.5;

Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.

SELECT title as titulo, rating, awards as premios FROM movies WHERE rating < 7.5;

Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.

SELECT title as titulo FROM movies ORDER BY title;

Mostrar los títulos de las primeras tres películas en la base de datos.

SELECT title as titulo FROM

Mostrar el top 5 de las películas con mayor rating.

SELECT title as titulo, rating FROM movies ORDER BY rating DESC LIMIT 5;

Mostrar las top 5 a 10 de las películas con mayor rating.

SELECT title as titulo, rating FROM movies ORDER BY rating DESC LIMIT 5 OFFSET 10;

Listar los primeros 10 actores (sería la página 1), 

SELECT first_name as nombre FROM actors LIMIT 10;

Luego usar offset para traer la página 3

SELECT first_name as nombre FROM actors LIMIT 20 OFFSET 30;


Hacer lo mismo para la página 5

SELECT first_name as nombre FROM actors LIMIT 40 OFFSET 50;

Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.



Mostrar a todos los actores cuyos nombres empiezan con Sam.

Mostrar el título de las películas que salieron entre el 2004 y 2008.

SELECT title as titulo 
FROM movies
WHERE release_date BETWEEN '2004-01-01'  AND '2008-12-31';


Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.

SELECT title as titulo, rating, release_date 
FROM movies
WHERE rating > 3 
AND awards > 1
AND release_date BETWEEN '1998-01-01'  AND '2009-12-31'
ORDER BY rating ASC;

Traer el top 3 a partir del registro 10 de la consulta anterior.

SELECT title as titulo, rating, release_date 
FROM movies
WHERE rating > 3 
AND awards > 1
AND release_date BETWEEN '1998-01-01'  AND '2009-12-31'
ORDER BY rating ASC
LIMIT 3 OFFSET 10;