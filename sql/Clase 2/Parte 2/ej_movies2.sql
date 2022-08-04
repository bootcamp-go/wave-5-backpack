-- 1
SELECT s.title, g.name 
FROM series s 
JOIN genres g 
WHERE s.genre_id = g.id;
-- 2
SELECT ep.title, a.first_name, a.last_name
FROM episodes ep 
JOIN actor_episode acp
JOIN actors a
WHERE acp.episode_id = ep.id
AND acp.actor_id = a.id;
-- 3
SELECT s.title, count(ss.serie_id) 
FROM series s
JOIN seasons ss
WHERE s.id = ss.serie_id
GROUP BY ss.serie_id;
-- 4
SELECT g.name, count(m.id) as Count
from genres g 
JOIN movies m
WHERE g.id = m.genre_id
GROUP BY m.genre_id
HAVING Count > 3;
-- 5
SELECT DISTINCT m.title, a.first_name, a.last_name
FROM movies m 
JOIN actor_movie acm
JOIN actors a
WHERE acm.movie_id = m.id
AND acm.actor_id = a.id
AND m.title LIKE 'La Guerra%'
ORDER BY m.title
