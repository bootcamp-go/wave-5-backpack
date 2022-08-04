USE movies_db;

SELECT ac.first_name nombre, ac.last_name apellido, mo.title titulo
FROM actors ac
INNER JOIN actor_movie mo_ac ON mo_ac.actor_id = ac.id
INNER JOIN movies mo ON mo_ac.movie_id = mo.id
WHERE mo.title LIKE '%Guerra de las galaxias%';