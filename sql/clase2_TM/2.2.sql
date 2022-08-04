USE movies_db;

SELECT ep.title titulo, ac.first_name nombre, ac.last_name apellido  
FROM episodes ep
INNER JOIN actor_episode ac_ep ON ep.id = ac_ep.episode_id
INNER JOIN actors ac ON ac.id = ac_ep.actor_id;