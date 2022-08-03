--1. Mostrar todos los registros de la tabla de movies.
SELECT * FROM movies;

--2. Mostrar el nombre, apellido y rating de todos los actores.
SELECT first_name AS name, last_name AS apellido, rating FROM actors;

--3. Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español
SELECT title AS titulo FROM series;

--4. Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
SELECT first_name, last_name FROM actors WHERE rating > 7.5;

--5. Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.
SELECT title, rating, awards FROM movies
WHERE rating > 7.5 AND awards > 2;
