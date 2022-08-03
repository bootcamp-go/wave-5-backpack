/* 2 */
select * from movies_db.movies;
/* 3 */
select first_name, last_name, rating FROM movies_db.actors ;
/* 4 */
select title Titulo from series;

/* 5 */
select first_name Nombre, LAST_NAME Apellido FROM MOVIES_DB.ACTORS WHERE RATING > 7.5 ;

/* 6 Mostrar el título de las películas, el rating y los premios de las películas con 
un rating mayor a 7.5 y con más de dos premios.*/
SELECT title, rating, awards from movies_db.movies WHERE RATING > 7.5 AND AWARDS > 2;
/* 7 Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente. */
select title, rating from movies order by rating;
/* 8 Mostrar los títulos de las primeras tres películas en la base de datos.*/
select title from movies limit 3;
/* 9 Mostrar el top 5 de las películas con mayor rating.*/
select title from movies order by rating desc limit 5;
/* 10 Mostrar las top 5 a 10 de las películas con mayor rating.*/
select title from movies order by rating desc limit 6 offset 4;
/* 11 Listar los primeros 10 actores (sería la página 1), */
select * from movies_db.actors limit 10;
/* 12 Luego usar offset para traer la página 3*/
select * from movies_db.actors limit 10 offset 29;

/* 13 Hacer lo mismo para la página 5 */
select * from movies_db.actors limit 10 offset 49;

/* 14 Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.*/
select title, rating from movies where title like '%toy story%';
/* 15 Mostrar a todos los actores cuyos nombres empiezan con Sam.*/
select * from actors where first_name like 'Sam%';
/* 16 Mostrar el título de las películas que salieron entre el 2004 y 2008. */
select title, release_date from movies where release_date between '20040101' and '20081231';
/* 17 Traer el título de las películas con el rating mayor a 3, con más de 1 premio y 
con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.*/
select title from movies where rating > 3 and awards > 1 and release_date between '19880101' and '20091231' order by rating desc;
/* 18 Traer el top 3 a partir del registro 10 de la consulta anterior */
select title from movies where rating > 3 and awards > 1 and release_date between '19880101' and '20091231' order by rating desc limit 3 offset 9;
