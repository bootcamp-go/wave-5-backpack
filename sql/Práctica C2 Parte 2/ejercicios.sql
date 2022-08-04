-- Active: 1659473777000@@127.0.0.1@3306@movies_db

-- 1. Explicar el concepto de normalización y para que se utiliza.

-- Es un proceso de validación de datos que se utiliza para eliminar las inconsistencias

-- o redundancias que se pueden generar en las tablas de la base de datos protegiendo la integridad.

-- 2. Agregar una película a la tabla movies.

INSERT INTO
    movies (
        created_at,
        updated_at,
        title,
        rating,
        awards,
        release_date,
        length
    )
VALUES (
        NOW(),
        NOW(),
        'The Shawshank Redemption',
        9.3,
        2,
        '1994-10-14',
        142
    );

-- 3. Agregar un género a la tabla genres.

INSERT INTO
    genres (
        created_at,
        updated_at,
        name,
        ranking,
        active
    )
VALUES (NOW(), NOW(), 'Anime', 13, 1);

-- 4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.

UPDATE movies
SET genre_id = (
        SELECT id
        FROM genres
        WHERE name = 'Anime'
    )
WHERE
    title = 'The Shawshank Redemption';

-- 5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.

UPDATE actors
SET favorite_movie_id = (
        SELECT id
        FROM movies
        WHERE
            title = 'The Shawshank Redemption'
    )
LIMIT 1;

-- 6. Crear una tabla temporal copia de la tabla movies.

CREATE TEMPORARY TABLE movies_copy ( SELECT * FROM movies );

-- 7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

DELETE FROM movies_copy WHERE awards < 5;

-- 8. Obtener la lista de todos los géneros que tengan al menos una película.

SELECT
    genres.name,
    #COUNT(movies.id) as movies_count
FROM genres
    JOIN movies ON genres.id = movies.genre_id
GROUP BY genres.name
HAVING COUNT(movies.id) > 0;

-- 9. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.

SELECT
    concat_ws(
        ' ',
        actors.first_name,
        actors.last_name
    ) as actor_name,
    #movies.title,
    #movies.awards
FROM actors
    JOIN movies ON actors.favorite_movie_id = movies.id
WHERE movies.awards > 3;

-- 10. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

EXPLAIN CREATE TEMPORARY TABLE movies_copy ( SELECT * FROM movies );

-- NO SE PUEDE HACER UN EXPLAIN CON CREATE

EXPLAIN DELETE FROM movies_copy WHERE awards < 5;

-- Se hace un escaneo completo de todas las filas de la tabla movies_copy para eliminar las filas que cumplen con el criterio de búsqueda.

-- 11. ¿Qué son los índices? ¿Para qué sirven?

-- Es una herramienta que permite acelerar la ejecución de consultas, mejoran mcuho los tiempos de respuesta de Querys complejas.

-- 12. Crear un índice sobre el nombre en la tabla movies.

CREATE INDEX movies_name_idx ON movies (title);

-- 13. Chequee que el índice fue creado correctamente.

SHOW INDEX FROM movies;