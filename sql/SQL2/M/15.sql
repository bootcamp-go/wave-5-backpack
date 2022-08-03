/*Mostrar a todos los actores cuyos nombres empiezan con Sam.*/

SELECT first_name, last_name
FROM actors
WHERE first_name LIKE "%sam"