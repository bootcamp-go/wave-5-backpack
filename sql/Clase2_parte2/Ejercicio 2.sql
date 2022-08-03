USE movies_db;

# EJERCICIO 1
SELECT s.title as titulo, g.name as nombre FROM series s
INNER JOIN genres g
ON s.genre_id = g.id;

# EJERCICIO 2
SELECT e.title as titulo, CONCAT(a.first_name, ' ', a.last_name) as nombre_completo 
FROM episodes e
INNER JOIN actor_episode ae
ON ae.episode_id = e.id
INNER JOIN actors a
ON ae.actor_id = a.id;

# EJERCICIO 3
SELECT se.title as titulo, MAX(s.number) as total_temporadas FROM series se 
INNER JOIN seasons s 
ON s.serie_id = se.id
GROUP BY se.title;

# EJERCICIO 4
SELECT g.name as nombre, COUNT(*) FROM genres g
INNER JOIN movies m
ON m.genre_id = g.id
GROUP BY m.genre_id
HAVING COUNT(*) >= 3;

# EJERCICIO 5
SELECT DISTINCT CONCAT(a.first_name, ' ', a.last_name) as nombre_completo 
FROM actors a
INNER JOIN actor_movie am
ON am.actor_id = a.id
INNER JOIN movies m
ON am.movie_id = m.id
where title like 'La Guerra de las galaxias%';