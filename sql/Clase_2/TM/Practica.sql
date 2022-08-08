 -- Primera Parte

-- 1. Se denomina JOIN a la sentencia que genera una intersección, la cual puede ser de diferentes tipos y bajo ciertas condiciones, entre dos columnas de diferentes tablas.
-- 2.1 INNER JOIN: es aquella sentencia que genera la intersección entre dos columnas que comparten una condición en común.
-- 2.2 LEFT JOIN: es aquella sentencia que trae los datos de la columna izquierda y que por otra parte comparte una condición con la columna derecha. Si la columna izquierda
	-- no comparte ninguna condición con la columna derecha, entonces se devuelve el valor null.
-- 3. Se utiliza para agrupar datos de una columna, a partir de la verificación y cumplimientos de ciertas condiciones.
-- 4. Se utiliza para hacer una filtración de los datos obtenidos a partir de un Group By.alter
-- 5.1 Corresponde a un inner join.
-- 5.2 Corresponde a un left join.
-- 6.1 

-- INER JOIN
USE movies_db;
SELECT mo.*, act.first_name, act.last_name, favorite_movie_id
	FROM movies mo
    INNER JOIN actors act 
    ON mo.id=act.favorite_movie_id;
    
-- RIGHT JOIN

SELECT mo.id, mo.title, mo.rating, mo.awards, act.first_name, act.last_name, favorite_movie_id
	FROM movies mo
    RIGHT JOIN actors act 
    ON mo.id=act.favorite_movie_id;

-- FULL JOIN

SELECT mo.id, mo.title, mo.rating, mo.awards, act.first_name, act.last_name, favorite_movie_id 
	FROM movies mo
    FULL OUTER JOIN actors act -- NO FUNCIONA
	ON mo.id=act.favorite_movie_id;
    
-- FULL WITH UNION ALL

USE movies_db;
SELECT mo.*, act.first_name, act.last_name, favorite_movie_id
	FROM movies mo
    RIGHT JOIN actors act 
    ON mo.id=act.favorite_movie_id

UNION ALL

SELECT mo.*, act.first_name, act.last_name, favorite_movie_id
	FROM movies mo
    LEFT JOIN actors act 
    ON mo.id=act.favorite_movie_id;
    
    
-- Segunda Parte

-- Ejericio 1
SELECT se.title, genres.name
FROM series se
INNER JOIN genres 
ON se.genre_id= genres.id;

-- Ejercicio 2
SELECT ep.title as 'Title', act.first_name as 'Name', act.last_name as 'Last Name'
FROM episodes ep
INNER JOIN actor_episode acte 
ON ep.id= acte.episode_id
INNER JOIN actors act
ON acte.actor_id= act.id;

-- Ejercicio 3
SELECT s.title AS 'Title', count(*) AS 'Seasons'
FROM series s 
INNER JOIN seasons ss 
ON s.id= ss.serie_id
GROUP BY s.title;

-- Ejercicio 4
SELECT g.name AS 'Name', count(*) as Number_of_Movies
FROM genres g
INNER JOIN movies m 
ON g.id= m.genre_id
GROUP BY name
HAVING Number_of_Movies>=3;

-- Ejercicio 5
SELECT concat(a.first_name," ",a.last_name) AS 'Name of Actors Star Wars'
FROM actor_movie am
INNER JOIN movies m ON m.id=am.movie_id
INNER JOIN actors a ON a.id=am.actor_id
WHERE m.title LIKE '%guerra de las galaxias%'
GROUP BY a.first_name,a.last_name;






    
    




