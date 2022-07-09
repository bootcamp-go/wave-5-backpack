#Mostrar el top 5 de las películas con mayor rating.
SELECT title titulo, rating FROM movies peliculas
ORDER BY rating DESC
LIMIT 10;