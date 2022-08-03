-- 2
SELECT * 
FROM movies;

-- 3
SELECT first_name Nombre, last_name Apellido, rating Rating 
FROM actors;

-- 4
SELECT title Titulo
FROM series AS Series;

-- 5
SELECT first_name Nombre, last_name Apellido
FROM actors
WHERE rating > 7.5;

-- 6
SELECT title Titulo, rating Rating, awards Premios
FROM movies
WHERE rating > 7.5 AND awards > 2;

-- 7
SELECT title Titulo, rating Rating
FROM movies
ORDER BY rating;

-- 8
SELECT title Titulo
FROM movies
LIMIT 3;

-- 9
SELECT *
FROM movies
ORDER BY rating DESC
LIMIT 5;

-- 10
SELECT *
FROM movies
ORDER BY rating DESC
LIMIT 5 OFFSET 5; 

-- 11
SELECT *
FROM actors
LIMIT 10;

-- 12
SELECT *
FROM actors
LIMIT 10 OFFSET 20;

-- 13
SELECT *
FROM actors
LIMIT 10 OFFSET 40;

-- 14
SELECT title Titulo, rating Rating
FROM movies
WHERE title LIKE "Toy Story%"
ORDER BY rating DESC;

-- 15
SELECT *
FROM actors
WHERE first_name LIKE "Sam%";

-- 16
SELECT title Titulo
FROM movies
WHERE release_date BETWEEN "2004-01-01" AND "2008-12-31";

-- 17
SELECT title Titulo
FROM movies
WHERE rating > 3 AND awards > 1 AND (release_date BETWEEN "1988-01-01" AND "2009-12-31")
ORDER BY rating DESC;

-- 18
SELECT title Titulo
FROM movies
WHERE rating > 3 AND awards > 1 AND (release_date BETWEEN "1988-01-01" AND "2009-12-31")
ORDER BY rating DESC
LIMIT 3 OFFSET 9