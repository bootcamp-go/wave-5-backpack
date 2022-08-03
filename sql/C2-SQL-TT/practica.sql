-- PRIMERA PARTE

-- 1. ¿A qué se denomina JOIN en una base de datos?

--    A la unión de 2 o mas tablas

-- 2. Nombre y explique 2 tipos de JOIN.

--    LEFT JOIN: Devuelve los resultados de la primera tabla, es decir la tabla izquieda
--    INNER JOIN: Devuelve los resultados que coinciden en 2 o mas tablas

-- 3. ¿Para qué se utiliza el GROUP BY?

--    Para agrupar el resultado según las columnas especificadas, lo que permite reducir el total de filas en el resultado

-- 4. ¿Para qué se utiliza el HAVING?

--    Es una condicional que afecta a los grupos obtenidos por GROUP BY

-- 5. Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:

--    El primero: corresponde a INNER JOIN
--    El segundo: corresponde a LEFT JOIN

-- 6. Escriba una consulta genérica por cada uno de los diagramas a continuación:

--    Esquema 1:

--      SELECT t1.*, t2.*
--      FROM TABLA1 t1
--      RIGHT JOIN
--      TABLA2 t2 ON t1.Id = t2.Id

--    Esquema 2:

--      SELECT * FROM t1
--      UNION
--      SELECT * FROM t2

-- SEGUNDA PARTE

-- 1. Mostrar el título y el nombre del género de todas las series.

SELECT s.title AS titulo_serie, g.name AS genero
FROM series s
INNER JOIN genres g ON s.genre_id = g.id;

-- 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

SELECT e.title AS titulo_episodio, a.first_name AS nombre, a.last_name AS apellido
FROM episodes e
INNER JOIN actor_episode ae ON e.id = ae.episode_id
INNER JOIN actors a ON ae.actor_id = a.id;

-- 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

SELECT s.title AS titulo_serie, COUNT(s2.number) AS total_temporadas
FROM series s
INNER JOIN seasons s2 ON s.id = s2.serie_id
GROUP BY s.title;

-- 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

SELECT g.name AS nombre_genero, COUNT(m.id) AS total_peliculas
FROM genres g
INNER JOIN movies m ON g.id = m.genre_id
GROUP BY g.name HAVING total_peliculas >= 3;

-- 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.

SELECT a.first_name AS nombre, a.last_name AS apellido
FROM actors a
INNER JOIN actor_movie am ON a.id = am.actor_id
INNER JOIN movies m ON am.movie_id = m.id 
WHERE m.title LIKE '%galaxias%'
GROUP BY a.first_name, a.last_name
