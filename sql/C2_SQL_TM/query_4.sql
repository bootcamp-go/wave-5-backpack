#Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, 
#siempre que sea mayor o igual a 3.
SELECT COUNT(*) as tot_pelis, g.name genero FROM movies mo
INNER JOIN genres g ON mo.genre_id = g.id
GROUP BY g.name HAVING tot_pelis >= 3 
ORDER BY tot_pelis ASC;

