/*Mostrar el título de las películas que salieron entre el 2004 y 2008*/
SELECT title, release_date
FROM movies
WHERE YEAR(release_date) BETWEEN 2004 AND 2088