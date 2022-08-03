-- PRIMERA PARTE

-- 1) ¿A qué se denomina JOIN en una base de datos?
-- R: Un JOIN es la unión de dos o más tablas mediante uno o varios campos en común.

-- 2) Nombre y explique 2 tipos de JOIN.
-- R: LEFT JOIN retorna todas las filas de la tabla de la izquierda y las filas coincidentes de la tabla de la derecha.
-- R: INNER JOIN Devuelve todas las filas cuando hay al menos una coincidencia en ambas tablas.

-- 3) ¿Para qué se utiliza el GROUP BY?
-- R: Para agrupar los resultados de una consulta, según los campos especificados.

-- 4) ¿Para qué se utiliza el HAVING?
-- R: Para agregar una condición a los resultados agrupados con GROUP BY.

-- 5) Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
-- R: INNER JOIN y LEFT JOIN

-- 6) Escriba una consulta genérica por cada uno de los diagramas a continuación:
-- R: SELECT tabla_A.*, tabla_B.* FROM tabla_A RIGHT JOIN tabla_B ON tabla_A.column_name == tabla_B.column_name
-- R: SELECT tabla_A.*, tabla_B.* FROM tabla_A FULL JOIN tabla_B ON tabla_A.column_name == tabla_B.column_name


-- SEGUNDA PARTE

-- Mostrar el título y el nombre del género de todas las series.
SELECT series.title as Serie, genres.name as Género
FROM series JOIN genres
ON series.genre_id = genres.id;

-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT e.title as Episodio, CONCAT(a.first_name, " ", a.last_name) as Actor
FROM episodes e JOIN actor_episode ae ON e.id = ae.episode_id
JOIN actors a ON a.id = ae.actor_id;

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT series.title as Serie, COUNT(seasons.id) as "Total Temporadas"
FROM series JOIN seasons ON series.id = seasons.serie_id
GROUP BY Serie;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT genres.name as Género, COUNT(movies.id) as Películas
FROM genres JOIN movies ON genres.id = movies.genre_id
GROUP BY Género
HAVING Películas >= 3;

-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT CONCAT(a.first_name, " ", a.last_name) as Reparto
FROM actors a JOIN actor_movie am ON a.id = am.actor_id
JOIN movies m ON m.id = am.movie_id
WHERE m.title LIKE "%Guerra de las galaxias%";
