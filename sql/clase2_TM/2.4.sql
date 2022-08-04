USE movies_db;

SELECT COUNT(*) total, ge.name genero
FROM genres ge
INNER JOIN movies mo ON mo.genre_id = ge.id
GROUP BY ge.name
HAVING total > 3;