/*Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.*/

SELECT s.title, COUNT(se.number)
FROM series s
LEFT JOIN seasons se ON s.id = se.serie_id
GROUP BY s.title