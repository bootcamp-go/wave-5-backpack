/*Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.*/

SELECT title, rating 
FROM movies
WHERE rating >7.5
AND awards >2