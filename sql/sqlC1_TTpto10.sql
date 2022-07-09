#Mostrar las top 5 a 10 de las películas con mayor rating.
SELECT title titulo, rating FROM movies peliculas
ORDER BY rating DESC
LIMIT 5 OFFSET 5;