select srs.title Titulo, gnr.name 
from series srs
join genres gnr
on srs.genre_id = gnr.id;

select ep.title 'Titulo Episodio', act.first_name Nombre, act.last_name Apellido
from episodes ep
join actor_episode ae 
on ae.episode_id = ep.id
join actors act
on ae.actor_id = act.id;

select sr.title Titulo, count(s.id) as Temporadas
from series sr
join seasons s
on s.serie_id = sr.id
group by sr.title
;

select gnr.name Genero, count(mv.id) as Movies
from genres gnr
join movies mv
on mv.genre_id = gnr.id
group by gnr.name
; 

select act.first_name Nombre, act.last_name
from actors act
join actor_movie am
on am.actor_id = act.id
join movies mv
on am.movie_id = mv.id
where mv.title like '%guerra%'
group by act.id
;





