SELECT * FROM actors a;

SELECT first_name, last_name, rating FROM actors a;

SELECT title titulo FROM series serie;

SELECT first_name, last_name, rating FROM actors a WHERE rating > 7.5;

SELECT title, rating, awards FROM movies m WHERE rating > 7.5 AND awards > 2;

 SELECT title, rating FROM movies m ORDER BY rating;
 
 SELECT id, title FROM movies m LIMIT 3;
 
 SELECT * FROM movies m ORDER BY rating DESC LIMIT 5;
 
 SELECT * FROM movies m ORDER BY rating DESC LIMIT 6 OFFSET 5;
 
 SELECT * FROM actors a LIMIT 10;
 
 SELECT * FROM actors a LIMIT 10 OFFSET 20;
 
 SELECT * FROM actors a LIMIT 10 OFFSET 40;
 
SELECT title, rating FROM movies m WHERE title = "Toy Story";

SELECT * FROM actors a WHERE first_name LIKE "Sam%";

SELECT title FROM movies m WHERE release_date BETWEEN '20040101' AND '20081231';

SELECT title, rating, awards, release_date 
FROM movies movieTop
WHERE rating > 3 AND awards > 1 AND release_date BETWEEN '19980101' AND '20091231' ORDER BY rating;

SELECT title, rating, awards, release_date 
FROM movies movieTop
WHERE rating > 3 AND awards > 1 AND release_date BETWEEN '19980101' AND '20091231' ORDER BY rating DESC LIMIT 3;

