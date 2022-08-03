USE movies_db;
SELECT title titulo, rating, awards premios 
FROM movies 
WHERE awards > 1 AND release_date BETWEEN '1998-01-01 00:00:00' AND '2009-01-01 00:00:00'
ORDER BY rating;