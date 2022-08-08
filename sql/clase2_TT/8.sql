USE movies_db;

SET sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));

SELECT name nombre, COUNT(*) peliculas
FROM genres ge
INNER JOIN movies mo ON mo.genre_id = ge.id
GROUP BY ge.name
HAVING peliculas >= 1;