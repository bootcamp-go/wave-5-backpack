/*
4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.
*/

UPDATE movies
SET genre_id = 13
WHERE id = 24;

UPDATE genres
SET created_at = "2022-07-07 22:00:00"
WHERE id = 13;

SELECT m.title, m.rating, m.awards, g.name
FROM movies AS m
LEFT JOIN genres AS g ON g.id = m.genre_id;
