#Mostrar el título de las películas que salieron entre el 2004 y 2008.
SELECT title, release_date FROM movies
WHERE release_date BETWEEN '2004-01-01 00:00:00' AND '2008-12-31 23:59:59';