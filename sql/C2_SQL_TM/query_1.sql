#Mostrar el título y el nombre del género de todas las series.
SELECT ser.title titulo, gen.name genero FROM series ser
INNER JOIN genres gen
ON ser.genre_id = gen.id;
