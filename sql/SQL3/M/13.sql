/*
Chequee que el índice fue creado correctamente
*/

-- CREATE INDEX title_index ON movies(title);
SHOW INDEX FROM movies;

SELECT *
FROM movies