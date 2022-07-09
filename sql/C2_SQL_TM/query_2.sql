#Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

SELECT ac.id id_actor, ac.first_name nombre, ac.last_name apellido, ep.title titulo_episodio FROM actors ac
INNER JOIN episodes ep
ON ep.id = ac.id;