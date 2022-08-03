select 
	e.title, 
	a.first_name,
	a.last_name 
from movies_db.episodes e 
join movies_db.actor_episode ae on ae.episode_id = e.id 
join movies_db.actors a on a.id = ae.actor_id