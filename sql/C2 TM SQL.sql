/*-----------------------PRIMERA PARTE-----------------------*/
/*¿A qué se denomina JOIN en una base de datos?
La sentencia JOIN permite combinar registros de una o mas tablas*/

/*Nombre y explique 2 tipos de JOIN.
LEFT JOIN: Trae todos los registros que tienen en comun ambas tablas Y todos los registros de la tabla izq 
RIGHT JOIN: Trae todos los registros que tienen en comun ambas tablas Y todos los registros de la tabla der 
*/

/*¿Para qué se utiliza el GROUP BY?
Se utiliza para agrupar resultados de los registros en base a una columna*/

/*¿Para qué se utiliza el HAVING?
Se utiliza de forma equivalente al Where pero es un filtro para los group by*/

/*Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
1) INNER JOIN
2) RIGHT JOIN*/

/*Escriba una consulta genérica por cada uno de los diagramas a continuación:*/
/*RIGHT JOIN*/
SELECT * FROM movies_db.series ser
RIGHT JOIN movies_db.genres gen
ON ser.genre_id = gen.id ;

/*INNER JOIN*/
SELECT * FROM movies_db.series ser, movies_db.genres gen
WHERE 1=1
AND ser.genre_id = gen.id ;


/*-----------------------SEGUNDA PARTE-----------------------*/
/*Mostrar el título y el nombre del género de todas las series.*/
SELECT ser.title titulo, gen.name genero
FROM movies_db.series ser, movies_db.genres gen
WHERE ser.genre_id = gen.id ;

/*Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.*/
SELECT episod.title, act.first_name, act.last_name
FROM movies_db.actor_episode actepisod, movies_db.actors act, movies_db.episodes episod
WHERE actepisod.actor_id = act.id
AND actepisod.episode_id = episod.id ;

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
FROM  movies_db.actors act,  movies_db.actor_movie act_mov, movies_db.movies mov
WHERE act.id = act_mov.actor_id
AND act_mov.movie_id = mov.id
AND mov.title LIKE 'La Guerra de las galaxias%'


