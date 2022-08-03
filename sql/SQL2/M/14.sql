/*Mostrar el título y rating de todas las películas cuyo título sea de Toy Story*/

SELECT title, rating
FROM movies
WHERE title LIKE "%toy story%"
