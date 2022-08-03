/*1*/
SELECT s.title, g.name 
FROM movies_db.series s 
INNER JOIN movies_db.genres g ON s.genre_id = g.id;

/*2*/
SELECT e.title, a.first_name, a.last_name 
FROM movies_db.episodes e 
JOIN movies_db.actor_episode ae ON e.id = ae.episode_id
JOIN movies_db.actors a ON ae.actor_id = a.id;

/*3*/
SELECT s.title, MAX(se.number) 
FROM movies_db.series s 
JOIN movies_db.seasons se ON s.id = se.serie_id
GROUP BY s.title;

/*4*/
SELECT g.name, COUNT(m.genre_id) 
FROM movies_db.genres g 
JOIN movies_db.movies m ON g.id = m.genre_id
GROUP BY m.genre_id
HAVING COUNT(m.genre_id) >= 3;

/*5*/
SELECT DISTINCT a.first_name, a.last_name
FROM  movies_db.actors a,  movies_db.actor_movie am, movies_db.movies m
WHERE a.id = am.actor_id
AND am.movie_id = m.id
AND m.title LIKE 'La Guerra de las galaxias%';
