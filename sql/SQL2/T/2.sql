/*Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.*/

SELECT e.title, a.first_name,a.last_name
FROM actors a
INNER JOIN actor_episode ac ON ac.actor_id = a.id
INNER JOIN episodes e ON ac.episode_id = e.id