USE movies_db;

-- PUNTO 2
SELECT * FROM movies;

-- PUNTO 3
SELECT first_name as nombre, last_name as apellido, rating as calificación FROM actors;

-- PUNTO 4
SELECT title as titulo FROM series as series;

-- PUNTO 5
SELECT first_name as nombre, last_name as apellido FROM actors WHERE rating > 7.5;

-- PUNTO 6
SELECT title as titulo, rating as calificacion, awards as premios FROM movies WHERE rating > 7.5 AND awards > 2;

-- PUNTO 7
SELECT title as titulo, rating as calificacion FROM movies ORDER BY rating;

-- PUNTO 8
SELECT title as titulo FROM movies LIMIT 3;

-- PUNTO 9
SELECT * FROM movies ORDER BY rating DESC LIMIT 5;

-- PUNTO 10
SELECT * FROM movies ORDER BY rating DESC LIMIT 6 OFFSET 4;

-- PUNTO 11
SELECT * FROM actors LIMIT 10;

-- PUNTO 12
SELECT * FROM actors LIMIT 10 OFFSET 20;

-- PUNTO 13
SELECT * FROM actors LIMIT 10 OFFSET 40;

-- PUNTO 14
SELECT title as titulo, rating as calificacion FROM movies WHERE title lIKE "Toy Story%";

-- PUNTO 15
SELECT * FROM actors WHERE first_name LIKE "Sam%";

-- PUNTO 16
SELECT title as titulo FROM movies WHERE release_date BETWEEN "2004-01-01" AND "2008-12-31";

-- PUNTO 17
SELECT title as titulo FROM movies WHERE rating > 3 AND awards > 1 AND release_date BETWEEN "1988-01-01" AND "2009-12-31" ORDER BY rating;

-- PUNTO 18
SELECT title as titulo FROM movies WHERE rating > 3 AND awards > 1 AND release_date BETWEEN "1988-01-01" AND "2009-12-31" ORDER BY rating DESC LIMIT 3 OFFSET 10;