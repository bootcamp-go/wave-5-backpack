#Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT COUNT(*) as cant_de_temporadas, se.title titulo FROM seasons s
INNER JOIN series se ON se.id = s.serie_id
GROUP BY se.title;
