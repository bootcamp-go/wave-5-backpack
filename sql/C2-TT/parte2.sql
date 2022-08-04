#PARTE 2

#Ej1
SELECT sr.title, gr.name FROM series sr
INNER JOIN genres gr ON sr.genre_id = gr.id;

#Ej2
SELECT ep.title, a.first_name, a.last_name
FROM episodes ep
INNER JOIN actor_episode ae ON ep.id = ae.episode_id
INNER JOIN actors a ON a.id = ae.actor_id;

#Ej3
SELECT title, COUNT(*) FROM episodes
GROUP BY title
ORDER BY title;

#Ej4
SELECT name, COUNT(*) FROM genres g
LEFT JOIN movies m ON g.id = m.genre_id
GROUP BY g.name
HAVING COUNT(*) > 3;

#Ej5
SELECT first_name, last_name FROM actors a
INNER JOIN actor_movie am ON am.actor_id = a.id
INNER JOIN movies m ON m.id = am.movie_id
WHERE m.title LIKE 'La Guerra de las galaxias%'
GROUP BY a.first_name, a.last_name;