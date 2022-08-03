##CLASE SQL 1 TT

##1.Mostrar todos los registros de la tabla de movies.

SELECT * FROM movies;

##3. Mostrar el nombre, apellido y rating de todos los actores.

SELECT first_name, last_name, rating FROM actors;

##4. Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español

SELECT title AS Titulo FROM series;

##5. Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.

SELECT first_name, last_name FROM actors WHERE rating > 7.5;

##6. Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.

SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2;

##7. Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.

SELECT title, rating FROM movies ORDER BY rating ASC;

##8. Mostrar los títulos de las primeras tres películas en la base de datos.

SELECT title FROM movies LIMIT 3 OFFSET 1;


##9.Mostrar el top 5 de las películas con mayor rating

SELECT title, rating  FROM movies ORDER BY rating DESC LIMIT 5;


##10.Mostrar las top 5 a 10 de las películas con mayor rating.

SELECT title, rating FROM movies ORDER BY rating DESC LIMIT 5 OFFSET 5;

##11.Listar los primeros 10 actores (sería la página 1), 

SELECT * FROM actors LIMIT 10;

##12.Luego usar offset para traer la página 3

SELECT * FROM actors LIMIT 10 OFFSET 20;

##13.Hacer lo mismo para la página 5

SELECT * FROM actors LIMIT 10 OFFSET 40;

##14.Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.

SELECT title, rating FROM movies WHERE title LIKE '%Toy Story%';

##15. Mostrar a todos los actores cuyos nombres empiezan con Sam.

SELECT * FROM actors WHERE first_name LIKE '%Sam%';

##16.Mostrar el título de las películas que salieron entre el 2004 y 2008.

SELECT title, release_date  FROM movies  WHERE YEAR(release_date) BETWEEN '2004' AND '2008';

##17.Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.

SELECT * FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN '1988' AND '2009' ORDER BY rating;


| id | created_at | updated_at | title                                               | rating | awards | release_date        | length | genre_id |
+----+------------+------------+-----------------------------------------------------+--------+--------+---------------------+--------+----------+
| 14 | NULL       | NULL       | Toy Story 2                                         |    3.2 |      2 | 2003-04-04 00:00:00 |    120 |        7 |
|  9 | NULL       | NULL       | Harry Potter y la cámara de los secretos            |    3.5 |      2 | 2009-08-04 00:00:00 |    200 |        8 |
| 11 | NULL       | NULL       | Alicia en el país de las maravillas                 |    5.7 |      2 | 2008-07-04 00:00:00 |    120 |     NULL |
| 19 | NULL       | NULL       | Big                                                 |    7.3 |      2 | 1988-02-04 00:00:00 |    130 |        8 |
|  2 | NULL       | NULL       | Titanic                                             |    7.7 |     11 | 1997-09-04 00:00:00 |    320 |        3 |
|  5 | NULL       | NULL       | Parque Jurasico                                     |    8.3 |      5 | 1999-01-04 00:00:00 |    270 |        5 |
| 12 | NULL       | NULL       | Buscando a Nemo                                     |    8.3 |      2 | 2000-07-04 00:00:00 |    110 |        7 |
| 15 | NULL       | NULL       | La vida es bella                                    |    8.3 |      5 | 1994-10-04 00:00:00 |   NULL |        3 |
|  4 | NULL       | NULL       | La Guerra de las galaxias: Episodio VII             |    9.0 |      6 | 2003-11-04 00:00:00 |    180 |        5 |
|  6 | NULL       | NULL       | Harry Potter y las Reliquias de la Muerte - Parte 2 |    9.0 |      2 | 2008-07-04 00:00:00 |    190 |        6 |
| 17 | NULL       | NULL       | Intensamente                                        |    9.0 |      2 | 2008-07-04 00:00:00 |    120 |        7 |
| 20 | NULL       | NULL       | I am Sam                                            |    9.0 |      4 | 1999-03-04 00:00:00 |    130 |        3 |
|  3 | NULL       | NULL       | La Guerra de las galaxias: Episodio VI              |    9.1 |      7 | 2004-07-04 00:00:00 |   NULL |        5 |
| 10 | NULL       | NULL       | El rey león                                         |    9.1 |      3 | 2000-02-04 00:00:00 |   NULL |       10


##18.Traer el top 3 a partir del registro 10 de la consulta anterior.

SELECT * FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN '1988' AND '2009' ORDER BY rating LIMIT 3 OFFSET 10;
