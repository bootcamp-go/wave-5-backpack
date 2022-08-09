/*
Obtener la lista de todos los géneros que tengan al menos una película.
*/

SELECT DISTINCT g.name
FROM genres AS g
INNER JOIN movies AS m ON m.genre_id = g.id