USE movies_db;

SELECT * FROM actors;

SELECT 	first_name,
		last_name,
		rating 
FROM actors;

SELECT programas.title AS titulo
FROM series AS programas; 

SELECT first_name 
FROM actors
where rating > 7.5;

SELECT title, rating, awards
FROM movies
where rating > 7.5 AND awards > 2;

SELECT title, rating
FROM movies
ORDER BY rating;

SELECT * FROM movies
LIMIT 3;

SELECT * FROM movies
ORDER BY rating DESC
LIMIT 5;

SELECT * FROM movies
ORDER BY rating DESC
LIMIT 5 OFFSET 5;

SELECT * FROM actors 
LIMIT 10;

SELECT * FROM actors 
LIMIT 10 OFFSET 20;

SELECT * FROM actors 
LIMIT 10 OFFSET 40;

SELECT * FROM movies
WHERE title LIKE '%Toy Story%';

SELECT * FROM actors
WHERE first_name LIKE 'Sam%';

SELECT title
FROM movies
WHERE rating > 3
AND awards > 1 
AND release_date BETWEEN '19880101' AND '20090101'
ORDER BY rating
LIMIT 3 OFFSET 10;







