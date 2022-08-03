-- Active: 1659473777000@@127.0.0.1@3306@movies_db

-- 2. Mostrar todos los registros de la tabla de movies.
SELECT *
FROM movies;
-- 3. Mostrar el nombre, apellido y rating de todos los actores.
SELECT first_name, last_name, rating
FROM actors;
-- 4. Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español
SELECT Novelas.title as Título
FROM series as Novelas;
-- 5. Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
-- 6. Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.
-- 7. Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
-- 8. Mostrar los títulos de las primeras tres películas en la base de datos.
-- 9. Mostrar el top 5 de las películas con mayor rating.
-- 10. Mostrar las top 5 a 10 de las películas con mayor rating.
-- 11. Listar los primeros 10 actores (sería la página 1), 
-- 12. Luego usar offset para traer la página 3
-- 13. Hacer lo mismo para la página 5
-- 14. Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
-- 15. Mostrar a todos los actores cuyos nombres empiezan con Sam.
-- 16. Mostrar el título de las películas que salieron entre el 2004 y 2008.
-- 17. Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.
-- 18. Traer el top 3 a partir del registro 10 de la consulta anterior.