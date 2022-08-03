USE movies_db;
SELECT * FROM movies;
SELECT first_name, last_name, rating FROM actors;
SELECT title AS titulo FROM series AS series_espanol;
SELECT first_name, last_name FROM actors
WHERE rating > 7.5; 
SELECT title, rating, awards FROM movies
WHERE rating > 7.5 AND awards > 2;
SELECT title, rating FROM movies
ORDER BY rating;
SELECT title FROM movies
LIMIT 3;
SELECT title, rating FROM movies
ORDER BY rating DESC
LIMIT 5;
SELECT title, rating FROM movies
ORDER BY rating DESC
LIMIT 5 OFFSET 6;
SELECT first_name, last_name FROM actors
LIMIT 10;
SELECT first_name, last_name FROM actors
LIMIT 10 OFFSET 20;
SELECT first_name, last_name FROM actors
LIMIT 10 OFFSET 40;
SELECT title, rating FROM movies
WHERE title = 'Toy Story';
SELECT first_name FROM actors
WHERE first_name = 'SAM';
SELECT title, release_date FROM movies
WHERE release_date BETWEEN '2004/01/01' AND '2008/01/01';
SELECT title, rating, awards, release_date FROM movies
WHERE rating > 3 AND awards > 1 AND release_date BETWEEN '1988/01/01' AND '2009/01/01'
ORDER BY rating;
SELECT title, rating, awards, release_date FROM movies
WHERE rating > 3 AND awards > 1 AND release_date BETWEEN '1988/01/01' AND '2009/01/01'
ORDER BY rating DESC LIMIT 3;