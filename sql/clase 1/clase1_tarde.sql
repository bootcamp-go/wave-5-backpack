#mostrar registro movies
select * from movies;
#Mostrar el nombre, apellido y rating de todos los actores.
select first_name,last_name,rating from actors;
#Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla 
#como el campo estén en español
select title as titulo from series;
#Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
select first_name,last_name,rating from actors where rating>7.5;
#Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 
#y con más de dos premios.
select title,rating,awards from movies where rating>7.5;
#Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
select title,rating from movies order by rating asc;
#Mostrar los títulos de las primeras tres películas en la base de datos.
select title from movies limit 3;
#Mostrar el top 5 de las películas con mayor rating.
select title,rating from movies order by rating desc limit 3;
#Mostrar las top 5 a 10 de las películas con mayor rating.
select title,rating from movies order by rating desc limit 5,5;
#Listar los primeros 10 actores
select * from actors limit 10;
#Luego usar offset para traer la página 3
SET @PageSize = 10;
select * from actors limit 10 offset 20;
#Hacer lo mismo para la página 5
select * from actors limit 10 offset 40;
#Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
select title,rating from movies where title LIKE '%Toy Story%';
#Mostrar a todos los actores cuyos nombres empiezan con Sam.
select * from actors where first_name LIKE 'Sam%';
#Mostrar el título de las películas que salieron entre el 2004 y 2008.
select title,release_date from movies where release_date>='2004-01-01' and release_date<'2009-01-01';
#Traer el título de las películas con el rating mayor a 3, con más de 1 premio y 
#con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.
#DECLARE @result1 TABLE (titulo string, ranking INT,premios int, estreno date);
select title,rating,awards,release_date from movies where rating>3 and awards>1 and release_date>='1988-01-01' and release_date<'2010-01-01' order by rating asc;
#Traer el top 3 a partir del registro 10 de la consulta anterior.
with tabla_anterior AS (select title,rating,awards,release_date from movies where rating>3 and awards>1 and release_date>='1988-01-01' and release_date<'2010-01-01' order by rating,awards asc)
 select * from tabla_anterior limit 10,3;