/*Agregar una película a la tabla movies.*/
INSERT INTO movies_db.movies VALUES (NULL,NULL,NULL,"SpiderMan",9.0,8,"2020-09-04 00:00:00",100,5);

/*Agregar un género a la tabla genres.*/
INSERT INTO movies_db.genres VALUES (NULL,"2020-09-04 00:00:00",NULL,"Policial",13,1) ;

/*Asociar a la película del Ej 2. con el género creado en el Ej. 3.*/
USE movies_db ;
UPDATE movies
SET genre_id = 13
WHERE movies.id = 22 ;

/*Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.*/
USE movies_db ;
UPDATE actors
SET favorite_movie_id = 22
WHERE actors.id = 1 ;

/*Crear una tabla temporal copia de la tabla movies.*/
USE movies_db ;
CREATE TEMPORARY TABLE movies_temp 
	(SELECT * FROM movies) ;
    
/*Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.*/
/*SET SQL_SAFE_UPDATES = 0;*/
DELETE FROM movies_db.movies_temp 
WHERE awards < 5 ;

/*Obtener la lista de todos los géneros que tengan al menos una película.*/
USE movies_db ;
SELECT name FROM genres
WHERE id IN (SELECT genre_id 
FROM movies
GROUP BY genre_id) ;

/*Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. */
USE movies_db ;
SELECT first_name, last_name
FROM actors act, movies mov
WHERE act.favorite_movie_id = mov.id
AND mov.awards > 3 ;


/*Crear un índice sobre el nombre en la tabla movies.*/
USE movies_db ;
CREATE INDEX movie_idx on movies(title) ;
/*Chequee que el índice fue creado correctamente.*/
SHOW INDEX FROM MOVIES ;