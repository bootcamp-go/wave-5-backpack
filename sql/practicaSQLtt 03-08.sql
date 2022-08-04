-- Mostrar el título y el nombre del género de todas las series.
Select series.title, genres.name from series join genres on series.genre_id = genres.id;

-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
Select episodes.title, actors.first_name, actors.last_name from episodes join actor_episode on episodes.id = actor_episode.episode_id
join actors on actors.id = actor_episode.actor_id;


-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
Select series.title, count(seasons.id) from series right join seasons on seasons.serie_id = series.id group by series.title;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
Select genres.name, count(movies.id) from movies left join genres on movies.genre_id = genres.id group by genres.name having count(movies.id) >= 3;


-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
Select actors.first_name, actors.last_name from actors join actor_movie on actors.id = actor_movie.actor_id
join movies on movies.id = actor_movie.movie_id where title like "La guerra de las galaxias%" group by actors.first_name, actors.last_name;

