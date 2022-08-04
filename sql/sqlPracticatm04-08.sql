-- Agregar una película a la tabla movies.
INSERT INTO movies (title, rating, awards, release_date, length, genre_id)
values ("Joker", 8.5, 1, "2019-02-02", 300, 3);

-- Agregar un género a la tabla genres.
INSERT INTO genres (name, ranking, active)
values ("otro", 13, 1);

-- Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies set genre_id = 13 where movies.title = "Joker";

-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors set favorite_movie_id = 13 where actors.id = 3;

-- 6 - Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_temp (select * from movies);
select * from movies_temp;

-- 7 - Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
delete from movies_temp where awards <5;

-- Obtener la lista de todos los géneros que tengan al menos una película.
select distinct genres.name from genres join movies on genres.id = movies.genre_id
where movies.genre_id is not null;

-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
select actors.first_name, actors.last_name, movies.title, movies.awards from actors join movies on actors.favorite_movie_id = movies.id
where movies.awards > 3;


-- Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
CREATE TEMPORARY TABLE movies_temp (select * from movies);
EXPLAIN delete from movies_temp where awards <5;


-- Crear un índice sobre el nombre en la tabla movies.
Create index movie_title_idx on movies(title);


-- Chequee que el índice fue creado correctamente.
show index from movies;