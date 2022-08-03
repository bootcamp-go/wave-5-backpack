/*Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, 
siempre que sea mayor o igual a 3.*/

SELECT g.name, COUNT(m.title)
FROM genres g
LEFT JOIN movies m ON g.id = m.genre_id
GROUP BY g.name
HAVING COUNT(m.title) >=3