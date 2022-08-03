USE movies_db;
SELECT se.title, gn.name
FROM series se
INNER JOIN genres gn
ON se.id = gn.id;
SELECT ep.title, act.first_name, act.last_name
FROM episodes ep
INNER JOIN actors act
ON ep.id = act.id;
SELECT se.title, MAX(tm.number) as total_temp
FROM series se
INNER JOIN seasons tm
ON se.id = tm.serie_id
GROUP BY title;
SELECT gn.name,  SUM(mo.genre_id) as total
FROM genres gn
INNER JOIN movies mo
ON gn.id = mo.genre_id
GROUP BY name HAVING total >= 3;
SELECT ac.first_name, ac.last_name
FROM actors ac
	INNER JOIN actor_movie am
    ON ac.id = am.actor_id
	INNER JOIN movies mo
    ON am.movie_id = mo.id
    WHERE mo.title LIKE '%galaxias%'
    GROUP BY ac.first_name, ac.last_name;




