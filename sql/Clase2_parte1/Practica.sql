USE movies_db;

-- PUNTO 1
SELECT s.title as titulo, g.name as nombre FROM series s
INNER JOIN genres g
ON s.genre_id = g.id;

-- PUNTO 2
SELECT e.title as titulo ,concat(a.first_name, ' ', a.last_name) as "nombre completo"
FROM actors a
INNER JOIN actor_episode ae
ON a.id = ae.actor_id
INNER JOIN episodes e
ON e.id = ae.episode_id;

-- PUNTO 3
SELECT s.title as titulo, max(se.number) as total_temporadas FROM series s
INNER JOIN seasons se
ON se.serie_id = s.id
GROUP BY s.title;

-- PUNTO 4
SELECT g.name as nombre, count(*) as total_peliculas
FROM movies m
INNER JOIN genres g
ON m.genre_id = g.id
GROUP BY g.name
HAVING count(*) >= 3;

-- PUNTO 5
SELECT DISTINCT a.first_name as nombre, a.last_name as apellido
FROM actors a
INNER JOIN actor_movie am
ON a.id = am.actor_id
INNER JOIN movies m
ON m.id = am.movie_id
WHERE m.title LIKE "La Guerra de las galaxias%";