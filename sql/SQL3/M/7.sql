/*
7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
*/
SET SQL_SAFE_UPDATES = 0;

DELETE FROM consu
WHERE awards IN(1,2,3,4);

SELECT *
FROM consu
