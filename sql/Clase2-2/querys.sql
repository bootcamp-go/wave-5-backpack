USE movies_db;

SELECT s.title AS Titulo,
		g.name AS Genero
FROM series AS s
JOIN genres g
ON s.genre_id = g.Id;

SELECT e.title AS Titulo_del_Episodio,
	   a.first_name AS Actor_Nombre,
	   a.last_name AS Actor_Apellido
FROM actor_episode AS a_e
JOIN episodes AS e
ON a_e.episode_id = e.Id
JOIN actors AS a
ON a_e.actor_id = a.Id;

SELECT se.title AS Titulo,
	   count(*) AS Total_temporadas
FROM series AS se 
JOIN seasons AS sa
ON se.Id = sa.serie_id 
GROUP BY Titulo;

SELECT g.name AS Nombre_Genero,
	   count(*) AS Total_Pelicilas
FROM genres AS g 
JOIN movies AS m
ON g.Id = m.genre_id 
GROUP BY Nombre_Genero HAVING Total_Pelicilas > 3;

SELECT ANY_VALUE(a.first_name) AS Nombre_Actor,
	   ANY_VALUE(a.last_name) AS Apellido_Actor
FROM actor_movie a_m
JOIN actors a
ON a_m.actor_id = a.Id 
JOIN movies m
ON a_m.movie_id = m.Id 
WHERE m.Title like '%La Guerra de las galaxias%'
GROUP BY Nombre_Actor




 




