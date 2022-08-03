#Mostrar el título y el nombre del género de todas las series.
select title,name from series join genres on series.genre_id= genres.id;
#Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
select episodes.title,actors.first_name,actors.last_name from episodes left join actor_episode on episodes.id =actor_episode.episode_id left join actors on actor_episode.actor_id=actors.id;
#Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select series.title,MAX(seasons.number) as num_temporadas from series left join seasons on series.id=seasons.serie_id group by series.title;
#Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
select genres.name,count(movies.genre_id) as num_peli from movies right join genres on movies.genre_id=genres.id group by genres.name having num_peli>=3;
#Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y 
#que estos no se repitan.
select distinct actors.first_name, actors.last_name from movies left join actor_movie on movies.id=actor_movie.movie_id left join actors on actors.id=actor_movie.actor_id where title LIKE '%guerra de las galaxias%';
