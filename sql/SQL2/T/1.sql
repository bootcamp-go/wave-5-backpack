/*Mostrar el título y el nombre del género de todas las series*/

SELECT s.title, g.name
FROM series s
LEFT JOIN genres g ON g.id = s.genre_id