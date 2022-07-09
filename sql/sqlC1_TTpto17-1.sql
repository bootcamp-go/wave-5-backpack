#Traer el título de las películas con el rating mayor a 3, con más de 1 premio y 
#con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.

	SELECT * FROM movies
    WHERE rating > 3 AND awards > 1 AND release_date BETWEEN '1988-01-01' AND '2009-12-31'
    ORDER BY rating DESC;
    