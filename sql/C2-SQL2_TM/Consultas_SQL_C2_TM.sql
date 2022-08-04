-- Mostrar el título y el nombre del género de todas las series.

SELECT
	se.title, ge.name
FROM
	series se
		INNER JOIN
	genres ge ON se.genre_id = ge.id;
    
-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

SELECT e.title AS "Episodio", concat(a.first_name, " ", a.last_name) AS "Actor" FROM episodes AS e 
	INNER JOIN actor_episode AS a_ep ON a_ep.episode_id = e.id
    INNER JOIN actors AS a ON a_ep.actor_id = a.id;
    
            
-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

SELECT s.title, COUNT(t.id) AS "Numero_temporadas" 
	FROM series AS s 
	RIGHT JOIN seasons t ON t.serie_id = s.id     
	GROUP BY s.title;
    
-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

SELECT g.name, COUNT(m.id) FROM movies AS m 
	LEFT JOIN genres AS g ON m.genre_id = g.id 
    GROUP BY g.name
    HAVING COUNT(m.id) >= 3;
    
-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.

SELECT ac.first_name, ac.last_name FROM actors AS ac 
	INNER JOIN actor_movie AS am ON ac.id = am.actor_id
    INNER JOIN movies AS m ON m.id = am.movie_id
    WHERE m.title LIKE "La Guerra de las Galaxias%"
    GROUP BY ac.first_name, ac.last_name;
