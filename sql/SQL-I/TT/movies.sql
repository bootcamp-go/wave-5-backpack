¿A qué se denomina JOIN en una base de datos?
El JOIN sirve para juntar dos tablas a partir de una o varias condiciones

Nombre y explique 2 tipos de JOIN.
INNER JOIN 
LEFT JOIN
¿Para qué se utiliza el GROUP BY?
Se usa para agrupar en base a una columna, por ejemplo por nombre
¿Para qué se utiliza el HAVING?
Es cómo un WHERE para el GROUP BY
Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:

Escriba una consulta genérica por cada uno de los diagramas a continuación:
INNER JOIN 
LEFT JOIN

Mostrar el título y el nombre del género de todas las series.
SELECT se.title titulo, ge.name genero
FROM series se
	JOIN genres ge ON se.genre_id = ge.id;

Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT ep.title titulo, ac.first_name nombre_actor, ac.last_name apellido_actor
FROM episodes ep
	JOIN actor_episode ae ON ep.id = ae.episode_id
	JOIN actors ac ON ac.id = ae.actor_id;

Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT title titulo, count(season_id) temporadas
FROM episodes
GROUP BY titulo
ORDER BY temporadas DESC;

Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, 
siempre que sea mayor o igual a 3.
SELECT ge.name genero, count(mo.id) peliculas
FROM genres ge
	JOIN movies mo ON ge.id = mo.genre_id
GROUP BY ge.name
HAVING peliculas >= 3; 
Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.