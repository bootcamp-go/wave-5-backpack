/*Mostrar el título y el nombre del género de todas las series.*/
SELECT ser.title titulo, gen.name genero
FROM movies_db.series ser, movies_db.genres gen
WHERE ser.genre_id = gen.id ;

/*Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.*/
SELECT episod.title, act.first_name, act.last_name
FROM movies_db.actor_episode actep, movies_db.actors act, movies_db.episodes episod
WHERE actep.actor_id = act.id
AND actep.episode_id = episod.id ;

/*Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.*/
SELECT ser.title, MAX(seas.number) temporadas
FROM movies_db.series ser, movies_db.seasons seas
WHERE ser.id = seas.serie_id
GROUP BY ser.title ;


/*Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno,siempre que sea mayor o igual a 3.*/
SELECT gen.name, COUNT(mov.genre_id) cantPelis
FROM movies_db.movies mov, movies_db.genres gen
WHERE gen.id = mov.genre_id
GROUP BY mov.genre_id
HAVING cantPelis >= 3 ;

/*Mostrar sólo el nombre y apellido de los actores que trabajan en todas 
las películas de la guerra de las galaxias y que estos no se repitan.*/
SELECT DISTINCT act.first_name, act.last_name
FROM  movies_db.actors act, movies_db.actor_movie act_mov, movies_db.movies mov
WHERE act.id = act_mov.actor_id
AND act_mov.movie_id = mov.id
AND mov.title LIKE 'La Guerra de las galaxias%' ;


/*CON SUBQUERY*/
SELECT DISTINCT act.first_name, act.last_name
FROM  movies_db.actors act, movies_db.actor_movie act_mov
WHERE act.id = act_mov.actor_id
AND act.id IN (SELECT act_mov.actor_id FROM movies_db.actor_movie act_mov, movies_db.movies mov
WHERE act_mov.movie_id = mov.id
AND mov.title LIKE 'La Guerra de las galaxias%')