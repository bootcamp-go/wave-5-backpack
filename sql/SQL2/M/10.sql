/*Mostrar las top 5 a 10 de las películas con mayor rating.*/
SELECT title, rating 
FROM movies
ORDER BY rating DESC
LIMIT 5 OFFSET 5