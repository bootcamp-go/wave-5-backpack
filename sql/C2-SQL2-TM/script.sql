-- 1
SELECT s.title Titulo, g.name Genero
FROM series s
INNER JOIN genres g ON s.genre_id = g.id;

-- 2 
SELECT e.title Titulo, a.first_name Nombre, a.last_name Apellido
FROM episodes e
INNER JOIN actor_episode ae ON e.id = ae.episode_id
INNER JOIN actors a ON ae.actor_id = a.id;

-- 3
SELECT COUNT(*), s.title Titulo
FROM series s
INNER JOIN seasons se ON se.serie_id = s.id
GROUP BY s.title;

-- 4
SELECT g.name Genero, COUNT(*)
FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.name
HAVING COUNT(*) > 3;

-- 5
SELECT a.first_name Nombre, a.last_name Apellido
FROM movies m
INNER JOIN actor_movie am ON am.movie_id = m.id
INNER JOIN actors a ON am.actor_id = a.id
WHERE m.title LIKE "La Guerra de las galaxias%"
GROUP BY a.first_name, a.last_name





