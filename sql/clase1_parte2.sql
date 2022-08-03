#2- Mostrar todos los registros de la tabla de movies.
select * from movies;
#3- Mostrar el nombre, apellido y rating de todos los actores.
select first_name, last_name, rating from actors;
#4- Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español
select title as Titulo  from series;
#5- Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
select first_name, last_name from actors where rating > 7.5;
#6- Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.
select title, rating, awards from movies
where rating > 7.5 
and awards > 2;
#7- Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
select title, rating from movies
order by rating;
#8- Mostrar los títulos de las primeras tres películas en la base de datos.
select title, rating from movies limit 3;
#9- Mostrar el top 5 de las películas con mayor rating.
select title, rating from movies 
order by rating desc
limit 5;
#10- Mostrar las top 5 a 10 de las películas con mayor rating.
select title, rating from movies 
order by rating desc
limit 5 offset 5;
#11- Listar los primeros 10 actores (sería la página 1), 
select * from actors limit 10;
#12- Luego usar offset para traer la página 3
select * from actors limit 10 offset 20;
#13- Hacer lo mismo para la página 5
select * from actors limit 10 offset 40;
#14- Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
select * from movies where title = 'Toy Story';
#15- Mostrar a todos los actores cuyos nombres empiezan con Sam.
select * from actors where 
first_name like 'Sam%';
#16- Mostrar el título de las películas que salieron entre el 2004 y 2008.
select title from movies 
where release_date between '2004-01-01' and '2008-12-31';
#17- Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.
select title from movies 
where rating > 3
and awards > 1 
and release_date between '1988-01-01' and '2009-12-31'
order by rating;
#18- Traer el top 3 a partir del registro 10 de la consulta anterior.
select title from movies 
where rating > 3
and awards > 1 
and release_date between '1988-01-01' and '2009-12-31'
order by rating
limit 3 offset 10;



