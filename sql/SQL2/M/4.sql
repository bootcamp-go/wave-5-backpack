/*Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español*/

SELECT title, CASE WHEN title ='Game of Thrones' THEN 'Juego de Tronos'
WHEN title ='Supernatural' THEN 'Supernatural'
WHEN title ='The Walking Dead' THEN 'Los muertos vivientes' 
WHEN title ='Person of Interest' THEN 'Persona de interés'
WHEN title ='The Big Bang Theory' THEN 'La teoría del Big Bang'
WHEN title ='Breaking Bad' THEN 'Breaking Bad' END AS Titulo_Espanol
FROM series