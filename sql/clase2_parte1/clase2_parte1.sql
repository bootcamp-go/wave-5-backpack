USE movies_db;

#2
SELECT * FROM movies;

#3
SELECT first_name, last_name, rating FROM actors;

#4
SELECT title AS titulo FROM series AS series;

#5
SELECT first_name, last_name FROM actors WHERE rating > 7.5;

#6
SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2;

#7
SELECT title, rating FROM movies ORDER BY rating;

#8
SELECT title FROM movies LIMIT 3 OFFSET 0;

#9
SELECT title FROM movies ORDER BY rating DESC LIMIT 5 OFFSET 0;

#10
SELECT title FROM movies ORDER BY rating DESC LIMIT 5 OFFSET 5;

#11
SELECT * FROM actors LIMIT 10;

#12
SELECT * FROM actors LIMIT 10 OFFSET 20;

#13
SELECT * FROM actors LIMIT 10 OFFSET 40;

#14
SELECT title, rating FROM movies WHERE title LIKE "%Toy Story%" ;

#15
SELECT * FROM actors WHERE first_name LIKE "%Sam%";

#16
SELECT * FROM movies WHERE release_date > '2004-01-01' AND release_date < '2008-12-31';

#17
SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND release_date > '1988-01-01' AND release_date < '2009-12-31' ORDER BY rating; 

#18
SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND release_date > '1988-01-01' AND release_date < '2009-12-31' ORDER BY rating LIMIT 3 OFFSET 9; 