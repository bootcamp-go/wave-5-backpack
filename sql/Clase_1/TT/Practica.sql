SELECT * FROM movies_db.movies;

SELECT first_name,last_name, rating FROM movies_db.actors;

SELECT title Título FROM movies_db.series;

SELECT first_name,last_name FROM movies_db.actors WHERE rating>7.5;

SELECT title, rating, awards FROM movies_db.movies WHERE rating>7.5;

SELECT title, rating FROM movies_db.movies ORDER BY title, rating;

SELECT title FROM movies_db.movies LIMIT 3 OFFSET 0;

SELECT title,rating FROM movies_db.movies ORDER by rating DESC LIMIT 10;

SELECT title,rating FROM movies_db.movies ORDER by rating DESC LIMIT 5 OFFSET 5;

SELECT first_name,last_name FROM movies_db.actors LIMIT 10;

SELECT first_name,last_name FROM movies_db.actors LIMIT 10 OFFSET 20;

SELECT first_name,last_name FROM movies_db.actors LIMIT 10 OFFSET 40;

SELECT title, rating FROM movies_db.movies WHERE title LIKE '%Toy Story%';

SELECT first_name,last_name FROM movies_db.actors WHERE first_name LIKE 'Sam%';

SELECT title, rating, release_date FROM movies_db.movies WHERE  release_date BETWEEN '2004-01-01' AND '2008-12-31';

SELECT title, rating, awards, release_date FROM movies_db.movies WHERE  rating>3 AND awards>1 AND release_date BETWEEN '1988-01-01' AND '2009-12-31' ORDER BY rating;

SELECT title, rating, awards, release_date FROM movies_db.movies WHERE  rating>3 AND awards>1 AND release_date BETWEEN '1988-01-01' AND '2009-12-31' ORDER BY rating LIMIT 3 OFFSET 9;



