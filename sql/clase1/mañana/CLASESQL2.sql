SELECT * FROM movies;

SELECT first_name,last_name,rating FROM actors;

SELECT title as 'NOMBRE' from series;

SELECT first_name, last_name FROM actors where rating > 7.5;

SELECT title, rating,awards FROM movies WHERE rating > 7.5 and rating >2;

SELECT title, rating from movies order by rating;

SELECT * FROM movies order by rating desc LIMIT 5;

SELECT * FROM movies limit 6 offset 4;

SELECT * FROM actors LIMIT 10;

SELECT * FROM actors LIMIT 10 offset 20;

SELECT title,rating FROM movies WHERE title like '%Toy Story%';

SELECT * FROM actors WHERE first_name like 'Sam%';

SELECT title FROM movies WHERE YEAR(release_date) BETWEEN 2004 AND 2008;

SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN 1988 AND 2009 ORDER BY rating;

SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN 1988 AND 2009 ORDER BY rating LIMIT 3 OFFSET 10;

