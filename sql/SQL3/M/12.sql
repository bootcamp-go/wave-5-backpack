/*
12. Crear un índice sobre el nombre en la tabla movies
*/

CREATE INDEX title_index ON movies(title);
-- SHOW INDEX FROM movies;

SELECT *
FROM movies