-- Agregar una película a la tabla movies.

INSERT INTO movies
(title, rating, awards,release_date, length, genre_id )
VALUES
('Minions', 9.0, 4, '2002-09-04 00:00:00',100, 7);

-- Agregar un género a la tabla genres.

INSERT INTO genres
(created_at, name, ranking, active)
VALUES
('2022-08-04 22:00:00','Anime', 13, 1);

-- Asociar a la película del Ej 2. con el género creado en el Ej. 3.

UPDATE movies
SET genre_id = 14
WHERE title = 'Minions';

-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.

UPDATE actors
SET favorite_movie_id = 22
WHERE id = 47;

-- Crear una tabla temporal copia de la tabla movies.

CREATE TEMPORARY TABLE movies_temp(SELECT * FROM movies);
SELECT * FROM movies_temp;
DROP TABLE movies_temp; 

-- Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

SELECT * FROM movies_temp;
DELETE FROM movies_temp
WHERE awards < 5;
SELECT * FROM movies_temp;

-- Obtener la lista de todos los géneros que tengan al menos una película.

SELECT g.name FROM genres AS g
	INNER JOIN movies AS m ON g.id = m.genre_id
GROUP BY g.name;

-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 

SELECT concat(a.first_name, " ", a.last_name), m.awards FROM actors AS a
	INNER JOIN movies AS m ON a.favorite_movie_id = m.id
    WHERE m.awards > 3;

-- Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

EXPLAIN SELECT * FROM movies_temp;
EXPLAIN DELETE FROM movies_temp
WHERE awards < 5;

-- Crear un índice sobre el nombre en la tabla movies.

CREATE INDEX title_idx ON movies(title);

-- Chequee que el índice fue creado correctamente.

SHOW INDEX FROM movies;
