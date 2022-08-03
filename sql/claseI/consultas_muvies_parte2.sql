-- PRIMERA PARTE!!
-- 1. ¿A qué se denomina JOIN en una base de datos?
-- join se denomina a las posibles interacciones de conjuntos de datos, union interseccion 
-- 2. Nombre y explique 2 tipos de JOIN.
-- INNER JOIN es la operacion de interseccion de dos conjuntos de datos
-- LEFT JOIN es operaccion se encarga de devolver los datos que coinciden entre dos grupos de datos y los que no coinciden del grupo de datos A o izquierdo
-- 3. ¿Para qué se utiliza el GROUP BY?
-- se utiliza para agrupar filas que coinciden con el mismo valor en la columna seleccionada
-- 4. ¿Para qué se utiliza el HAVING?
-- lo utilizamos para realizar un condicional en un GROUP BY
-- 5. Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
-- ejemplo 1 INNER JOIN, ejemplo 2 LEFT JOIN
-- 6. Escriba una consulta genérica por cada uno de los diagramas a continuación:
-- ejemplo 1
SELECT column_x
FROM table1
RIGHT JOIN table2
ON table1.column_x = table2.column_x;
-- ejemplo 2
SELECT column_x
FROM table1
FULL JOIN table2
ON table1.column_x = table2.column_x;
-- SEGUNDA PARTE
USE movies_db;
-- 1. Mostrar el título y el nombre del género de todas las series.
SELECT s.title titulo, g.name genero
FROM series s
INNER JOIN 
	genres g ON s.genre_id = g.id;
-- 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT e.title titulo, a.first_name nombre, a.last_name apellido
FROM episodes e
INNER JOIN actor_episode ae ON ae.episode_id = e.id
INNER JOIN actors a ON ae.actor_id = a.id;
-- 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT s.title titulo, COUNT(*)
FROM series s
INNER JOIN seasons se ON s.id = se. serie_id
GROUP BY s.title;
-- 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT g.name, COUNT(*)
FROM movies m
INNER JOIN genres g ON m.genre_id = g.id
GROUP BY g.name;
-- 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT a.first_name nombre, a.last_name apellido
FROM movies m
INNER JOIN actor_movie am ON m.id = am.movie_id
INNER JOIN actors a ON am.actor_id = a.id
WHERE
	m.title LIKE 'La Guerra de las galaxias%';
