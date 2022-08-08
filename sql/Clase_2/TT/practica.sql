-- Ejercicio 1

/* Normalización consiste en estandarizar y validar los datos de una base de dato, de tal manera de evitar las redundancias o errores, protegiendo su
 integridad y mejorando la interpretación para que resulte más facil hacer las consultas y eficiente.
*/
-- Ejercicio 2

SELECT * FROM movies;

INSERT INTO movies
(title, rating, awards, release_date, length, genre_id)
VALUES
("La inventada",9.0,2,'2003-11-04 00:00:00',200,8);

-- Ejercicio 3

SELECT * FROM movies_db.genres;

INSERT INTO genres
(created_at,name,ranking,active);

-- Ejercicio 4

UPDATE movies SET genre_id=13 WHERE id=22;

-- Ejercicio 5

UPDATE actors SET favorite_movie_id=22 WHERE id=3;

-- Ejercicio 6

CREATE TEMPORARY TABLE copia_movies
	SELECT * FROM movies;

-- Ejercicio 7
SELECT * FROM copia_movies;

DELETE FROM copia_movies
WHERE awards<5;

-- Ejercicio 8

SELECT g.name as 'Género'
FROM genres g
INNER JOIN movies m
ON g.id=m.genre_id
GROUP BY g.name;

-- Ejercicio 9

SELECT a.first_name, a.last_name
FROM actors a
INNER JOIN movies m
ON a.favorite_movie_id=m.id
GROUP BY a.first_name, a.last_name, m.awards
HAVING m.awards>3;


-- Ejercicio 11

/* Un índice SQL es una tabla de búsqueda rápida para poder encontrar los registros que los usuarios necesitan buscar con mayor frecuencia. 
Ya que un índice es pequeño, rápido y optimizado para búsquedas rápidas. Además, que son muy útiles para conectar 
las tablas relacionales y la búsqueda de tablas grandes.
*/

-- Ejercicio 12

CREATE INDEX movies_title
	ON movies (title);
    
-- Ejercicio 13

SHOW INDEX FROM movies;

