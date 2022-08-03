/*Mostrar todos los registros de la tabla de movies*/
SELECT * FROM movies m;

/*Mostrar todos los registros de la tabla de movies*/
SELECT first_name, last_name, rating FROM actors a;

/*Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español*/
SELECT m.title as "Titulo" FROM movies m;

/*Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español*/
SELECT a.first_name, a.last_name FROM actors a WHERE a.rating>7.5;

/*Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.*/
SELECT m.title, m.rating, m.awards FROM movies m WHERE m.rating>7.5 AND m.awards>2;

/*Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.*/
SELECT m.title, m.rating FROM movies m ORDER BY m.rating;

/*Mostrar los títulos de las primeras tres películas en la base de datos.*/
SELECT m.title FROM movies m LIMIT 3;

/*Mostrar el top 5 de las películas con mayor rating.*/
SELECT m.title, m.rating FROM movies m ORDER BY m.rating DESC LIMIT 5;

/*Mostrar las top 5 a 10 de las películas con mayor rating.*/
SELECT m.title, m.rating FROM movies m ORDER BY m.rating DESC LIMIT 5 OFFSET 10;

/*Listar los primeros 10 actores (sería la página 1)*/
SELECT * FROM actors LIMIT 10;

/*Luego usar offset para traer la página 3*/
SELECT * FROM actors LIMIT 10 OFFSET 30;

/*Hacer lo mismo para la página 5(paginas de 5 datos)*/
SELECT * FROM actors LIMIT 5 OFFSET 25;

/*Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.*/
SELECT m.title, m.rating FROM movies m WHERE m.title LIKE '%Toy Story%';

/*Mostrar a todos los actores cuyos nombres empiezan con Sam.*/
SELECT * FROM actors AS a WHERE a.first_name LIKE 'Sam%';

/*Mostrar el título de las películas que salieron entre el 2004 y 2008.*/
SELECT m.title FROM movies m 
WHERE m.release_date > '2004-01-01 00:00:00'
AND m.release_date < '2008-12-31 00:00:00';

/*Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.*/
SELECT m.id, m.title, m.rating FROM movies m 
WHERE m.rating > 3
AND m.awards > 1
AND m.release_date BETWEEN '1988-01-01 00:00:00' AND '2009-12-31 00:00:00'
ORDER BY m.rating;

/*Traer el top 3 a partir del registro 10 de la consulta anterior.*/
SELECT m.title, m.rating FROM movies m 
WHERE m.rating > 3
AND m.awards > 1
AND m.release_date BETWEEN '1988-01-01 00:00:00' AND '2009-12-31 00:00:00'
ORDER BY m.rating
LIMIT 3 OFFSET 10;