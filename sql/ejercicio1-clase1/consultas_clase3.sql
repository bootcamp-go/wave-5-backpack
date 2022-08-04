insert into movies (title, rating, awards, release_date, length, genre_id)
values('Berserk: The Golden Age Arc I - Haou no Tamago',7.5, 2, '2012-02-04 00:00:00', 76, 7);
insert into movies (title, rating, awards, release_date, length, genre_id)
values('Berserk: The Golden Age Arc II - Dorudorei koryaku',7.7, 1, '2012-06-23 00:00:00', 95, 7);
insert into movies (title, rating, awards, release_date, length, genre_id)
values('Berserk: The Golden Age Arc III - Kourin',7.8, 2, '2013-02-01 00:00:00', 110, 7);

insert into genres (created_at, name, ranking, active)
values('2022-08-05 00:00:00', 'Anime', 13, 1);

UPDATE movies SET genre_id =13 WHERE id = 26;

UPDATE actors SET favorite_movie_id = 26 WHERE id = 12;

create temporary table temp_movies as select * from movies;

delete from temp_movies where awards < 5 ;

select gn.name from genres gn
join movies mv
on mv.genre_id = gn.id
group by gn.name;

select act.first_name, act.last_name from actors act
join movies mv
on act.favorite_movie_id = mv.id
where mv.awards > 3;

explain select gn.name from genres gn
join movies mv
on mv.genre_id = gn.id
group by gn.name;

explain select act.first_name, act.last_name from actors act
join movies mv
on act.favorite_movie_id = mv.id
where mv.awards > 3;

show index from movies;


select * from movies;
select * from genres;
select * from actors;
select * from temp_movies;
drop temporary table temp_movies;
