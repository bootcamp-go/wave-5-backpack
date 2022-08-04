#Agregar una película a la tabla movies.
INSERT INTO movies 
(title, rating, awards, release_date,length,genre_id)
VALUES 
('Charli y la fabrica de chocolates', 8.5, 4, '1999-01-04 00:00:00',180, 4);
select * from movies;
#Agregar un género a la tabla genres.
select * from genres;
INSERT INTO genres 
(created_at, name, ranking, active)
VALUES 
('2022-07-03 23:00:00', 'lugubre', 13, 1);
select * from genres;
#Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies 
SET genre_id = 13
WHERE id = 22 or id=23;
select * from movies;
#Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
select * from actors;
UPDATE actors 
SET favorite_movie_id = 22
WHERE id = 3;
select * from actors;
#Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_aux 
select * from movies;
select * from movies_aux;
#Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
DELETE FROM movies_aux WHERE awards < 5;
select * from movies_aux;
#Obtener la lista de todos los géneros que tengan al menos una película.
select distinct genres.name from movies left join genres on movies.genre_id = genres.id;
#Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
select actors.first_name,actors.last_name from actors left join movies on actors.favorite_movie_id = movies.id where movies.awards>3;
#Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
#explain CREATE TEMPORARY TABLE movies_aux select * from movies;
explain DELETE FROM movies_aux WHERE awards < 5;
# los indices sirven para agilizar ciertas consultas frecuentes, que apuntan con gran frecuencia a una columna en particular de la tabla, estonces lo que hace el indice es dejar a facil alcance aquella columna que se suele usar con tanta frecuencia
#Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX movies_title_idx
    ON MOVIES(title);
#Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;

