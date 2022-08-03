/* EJERCICIO 2*/
SELECT * 
FROM movies_db.movies;
/* EJERCICIO 3*/
SELECT first_name as nombre, last_name as apellido, rating
FROM movies_db.actors;
/* EJERCICIO 4*/
SELECT title as titulo 
FROM movies_db.series as series;

/* EJERCICIO 5*/
SELECT first_name as nombre, last_name as apellido 
FROM movies_db.actors 
WHERE rating > 7.5;
/* EJERCICIO 6*/
SELECT title as titulo, rating , awards as premios 
FROM movies_db.movies 
WHERE 
	rating > 7.5 and
	awards > 2;
/* EJERCICIO 7*/
SELECT title as titulo, rating , awards as premios 
FROM movies_db.movies 
ORDER BY rating ASC;
/* EJERCICIO 8*/
SELECT title as titulo 
FROM movies_db.movies 
LIMIT 3;
/* EJERCICIO 9*/
SELECT * 
FROM movies_db.movies 
ORDER BY rating DESC 
LIMIT 10;
/* EJERCICIO 10*/
SELECT * 
FROM movies_db.movies 
ORDER BY rating DESC
LIMIT 5 OFFSET 5;
/* EJERCICIO 11*/
SELECT * 
FROM movies_db.actors 
LIMIT 10;
/* EJERCICIO 12*/
SELECT * 
FROM movies_db.actors 
LIMIT 10 OFFSET 20;
/* EJERCICIO 13*/
SELECT * 
FROM movies_db.actors 
LIMIT 10 OFFSET 40;
/* EJERCICIO 14*/
SELECT title as titulo, rating 
FROM movies_db.movies as mv  
WHERE 
	mv.title = "Toy Story";
/* EJERCICIO 15*/
SELECT first_name as nombre 
FROM movies_db.actors 
WHERE 
	first_name like "Sam%";
/* EJERCICIO 16*/
SELECT title as titulo 
FROM movies_db.movies 
WHERE 
	DATE(release_date) BETWEEN '2004-00-00' AND '2008-12-30';
/* EJERCICIO 17*/
SELECT *
FROM movies_db.movies AS peliculas
WHERE 
	rating > 3 AND 
	awards > 1 AND 
	DATE(release_date) BETWEEN '1988-01-01' AND '2009-12-31'
ORDER BY rating DESC;
/* EJERCICIO 18*/
SELECT title as titulo
FROM movies_db.movies 
WHERE 
	rating > 3 AND 
	awards > 1 AND 
	DATE(release_date) BETWEEN '1988-00-00' AND '2009-12-30'
ORDER BY rating DESC
LIMIT 3 OFFSET 10;