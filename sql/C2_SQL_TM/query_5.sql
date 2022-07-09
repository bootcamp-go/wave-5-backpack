#Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias 
#y que estos no se repitan.
SELECT a.first_name nombre, a.last_name apellido FROM actors AS a
INNER JOIN actor_movie AS am ON a.id = am.actor_id
INNER JOIN movies AS m ON m.id = am.movie_id
WHERE m.title LIKE '%galaxias%'
GROUP BY a.first_name, a.last_name;