select * from movies;
select first_name Nombre, last_name Apellido from actors;
select title Titulo from series;
select first_name Nombre, last_name Apellido from actors where rating > 7.5;
select title Titulo, rating Rating, awards Premios from movies where rating > 7.5 and awards >= 2;
select title Titulo, rating Rating from movies order by  rating asc;
select title Titulo from movies limit 3;
select * from movies order by rating desc limit 5 ;
select * from movies order by rating desc limit 5 offset 2;
select * from actors limit 10 offset 3;
select * from actors limit 10 offset 5;
select title Titulo, rating Rating from movies where title like 'Toy Story';
select * from actors where first_name like 'Sam%';
select * from movies where release_date between '2001-01-01' and '2008-01-01';
select * from movies where rating > 3 and awards > 1 and release_date between '1988-01-01' and '2009-01-01' order by rating;
select * from movies where rating > 3 and awards > 1 and release_date between '1988-01-01' and '2009-01-01' order by rating desc limit 3;





