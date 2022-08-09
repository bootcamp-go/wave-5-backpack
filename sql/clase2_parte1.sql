USE movies_db;

-- Repaso funciones de agregacion

SELECT 
    COUNT(*) AS cantidad
FROM
    actors
WHERE
    rating = 7.5 AND favorite_movie_id = 1;

-- Inner Join

SELECT 
    mo.id, ac.favorite_movie_id, ac.first_name, ac.last_name
FROM
    movies mo
        INNER JOIN
    actors ac ON mo.id = ac.favorite_movie_id;

-- Left Join

SELECT 
    mo.id, ac.favorite_movie_id, mo.title, ac.id, ac.first_name
FROM
    movies mo
        LEFT JOIN
    actors ac ON mo.id = ac.favorite_movie_id;

-- Right Join

SELECT 
    mo.id, ac.favorite_movie_id, mo.title, ac.id, ac.first_name
FROM
    movies mo
        RIGHT JOIN
    actors ac ON mo.id = ac.favorite_movie_id;

-- Full Join ?

SELECT 
    *
FROM
    movies mo
        LEFT JOIN
    actors ac ON mo.id = ac.favorite_movie_id 
UNION 
SELECT 
    *
FROM
    movies mo
        RIGHT JOIN
    actors ac ON mo.id = ac.favorite_movie_id;

-- Antes de los siguientes ejemplos:

SET sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));

-- Group By

SELECT 
    COUNT(*), mo.title, mo.rating, mo.awards
FROM
    movies mo
        INNER JOIN
    actors ac ON mo.id = ac.favorite_movie_id
GROUP BY title;

-- Having

SELECT 
    COUNT(*) as tot_act, mo.title, mo.rating, mo.awards
FROM
    movies mo
        INNER JOIN
    actors ac ON mo.id = ac.favorite_movie_id
GROUP BY title
HAVING tot_act > 2;

-- Group by - having - where

SELECT 
    awards, COUNT(*)
FROM
    movies
WHERE
    rating > 8
GROUP BY awards
HAVING awards > 2
ORDER BY awards DESC;

-- Subqueries

SELECT 
    *
FROM
    actor_movie
WHERE
    movie_id IN (SELECT 
            id
        FROM
            movies
        WHERE
            rating = 9.0);
SELECT 
            id
        FROM
            movies
        WHERE
            rating = 9.0;