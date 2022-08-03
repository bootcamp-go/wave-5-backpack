-- 2
SELECT * from movies;
-- 3
SELECT a.first_name, a.last_name, a.rating 
FROM actors a;
-- 4
SELECT s.title Titulo 
from series s;
-- 5
SELECT a.first_name, a.last_name 
FROM actors a 
WHERE a.rating > 7.5;
-- 6
SELECT m.title, m.rating, m.awards 
FROM movies m 
WHERE m.rating > 7.5 
AND m.awards > 2;
-- 7
SELECT m.title, m.rating 
FROM movies m 
ORDER BY m.rating;
-- 8
SELECT m.title 
FROM movies m 
LIMIT 3;
-- 9
SELECT m.title, m.rating 
FROM movies m 
ORDER BY m.rating 
DESC LIMIT 5;
-- 10
SELECT m.title, m.rating 
FROM movies m 
ORDER BY m.rating DESC 
LIMIT 5 OFFSET 5;
-- 11
SELECT * FROM actors 
LIMIT 5;
-- 12 
SELECT * FROM actors 
LIMIT 5 OFFSET 10;
-- 13 
SELECT * FROM actors 
LIMIT 5 OFFSET 20;
-- 14 
SELECT m.title, m.rating 
FROM movies m 
WHERE m.title LIKE 'Toy Story%';
-- 15
SELECT * from actors a 
WHERE a.first_name LIKE 'Sam%';
-- 16
SELECT m.title 
FROM movies m 
WHERE m.release_date BETWEEN '2004-01-01 00:00:00' AND '2008-12-31 23:59:59';
-- 17
SELECT m.title 
FROM movies m 
WHERE m.rating > 3 
AND m.awards > 1 
AND m.release_date BETWEEN '1988-01-01 00:00:00' AND '2009-12-31 23:59:59' 
ORDER BY m.rating;
-- 18
SELECT m.title 
FROM movies m 
WHERE m.rating > 3 
AND m.awards > 1 
AND m.release_date BETWEEN '1988-01-01 00:00:00' AND '2009-12-31 23:59:59' 
ORDER BY m.rating
LIMIT 3 OFFSET 10;