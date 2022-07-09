#Traer el top 3 a partir del registro 10 de la consulta anterior.
SELECT * FROM movies
WHERE rating > 3 AND awards > 1 AND release_date BETWEEN '1988-01-01' AND '2009-12-31'
ORDER BY rating DESC
LIMIT 3 OFFSET 10;