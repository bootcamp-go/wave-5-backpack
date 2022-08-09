/*
6. Crear una tabla temporal copia de la tabla movies
*/

CREATE TEMPORARY TABLE consu AS (
SELECT title, rating, awards,release_date,length, genre_id
FROM movies);

SELECT *
FROM consu