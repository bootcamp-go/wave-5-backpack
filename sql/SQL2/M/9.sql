/*Mostrar el top 5 de las películas con mayor rating.*/
SELECT title, rating 
FROM movies
ORDER BY rating DESC
LIMIT 3