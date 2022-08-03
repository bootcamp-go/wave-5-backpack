/*Traer el top 3 a partir del registro 10 de la consulta anterior*/

SELECT title, rating
FROM movies
WHERE rating >3
AND awards >1
AND YEAR(release_date) BETWEEN 1988 AND 2009
ORDER BY rating DESC
LIMIT 3 OFFSET 10
