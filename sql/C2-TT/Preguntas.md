# C2-TT
## Preguntas y Respuestas

1. Explicar el concepto de normalización y para que se utiliza.
    La normalizacion es un proceso que sirve para limpiar la data de redundancias e inconsistencias en pro de la estandarizacion, integridad
    y favoreciendo su interpretacion.

2. Agregar una película a la tabla movies.
    INSERT INTO movies
    (title, rating, awards,release_date,length, genre_id)
    VALUES
    ('Rapido y Furioso: reto Tokio',8.5,1,'2006-10-04',104,4);

3. Agregar un género a la tabla genres.
    INSERT INTO genres
    (name, ranking, active)
    VALUES
    ('Autos',13,1);

4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.
    UPDATE movies
    SET genre_id = 13
    WHERE id = 22;

5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2
    UPDATE actors
    SET favorite_movie_id = 13
    WHERE id IN (4,21);

6. Crear una tabla temporal copia de la tabla movies.
    CREATE TEMPORARY TABLE movies_temp 
    SELECT * FROM movies;

7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
    DELETE FROM movies_temp WHERE awards < 5;

8. Obtener la lista de todos los géneros que tengan al menos una película.
    SELECT gn.name, COUNT(mo.genre_id) as total_movies
    FROM genres gn
    INNER JOIN movies mo
    ON gn.id = mo.genre_id
    GROUP BY gn.name HAVING COUNT(mo.genre_id)>1;

9. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
    SELECT ac.first_name, ac.last_name, mo.title, mo.awards
    FROM actors ac
    INNER JOIN movies mo
    ON ac.favorite_movie_id = mo.id
    GROUP BY ac.first_name, ac.last_name,mo.title, mo.awards HAVING mo.awards > 3; 

10. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
    6) Explain data not available for statement
    7) /query_block #1/
            /21 rows
        /Full Table Scan/
        movies_temp

11. ¿Qué son los índices? ¿Para qué sirven?
    Son espacios en memoria que nos permiten acceder a consultas SQL de manera mas rapida y eficiente

12. Crear un índice sobre el nombre en la tabla movies.
    CREATE INDEX movies_title_idx
    ON movies (title)
    
13. Chequee que el índice fue creado correctamente. 
    SHOW INDEX FROM movies;
