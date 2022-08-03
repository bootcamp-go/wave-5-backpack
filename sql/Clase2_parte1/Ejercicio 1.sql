USE movies_db;

# EJERCICIO 2
SELECT * FROM movies;

# EJERCICIO 3
SELECT first_name as nombre, last_name as apellido, rating as clasificacion FROM actors;

# EJERCICIO 4
SELECT title as titulo FROM series as series;

# EJERCICIO 5
SELECT first_name as nombre, last_name as apellido FROM actors WHERE rating > 7.5;

# EJERCICIO 6
SELECT title as titulo, rating as clasificacion, awards as premios FROM movies WHERE rating > 7.5 AND awards  > 2; 

# EJERCICIO 7
SELECT title as titulo, rating as clasificacion FROM movies ORDER BY rating ASC;

# EJERCICIO 8
SELECT title as titutlo FROM movies LIMIT 3;

# EJERCICIO 9
SELECT * FROM movies ORDER BY rating DESC LIMIT 5;

# EJERCICIO 10
SELECT * FROM movies ORDER BY rating DESC LIMIT 6 OFFSET 4;

# EJERCICIO 11
SELECT * FROM actors LIMIT 10;

# EJERCICIO 12
SELECT * FROM actors LIMIT 10 OFFSET 20;

# EJERCICIO 13
SELECT * FROM actors LIMIT 10 OFFSET 40;

# EJERCICIO 14
SELECT title as titulo, rating as clasificacion FROM movies WHERE title LIKE 'Toy Story%';

# EJERCICIO 15
SELECT * FROM actors WHERE first_name LIKE 'Sam%';

# EJERCICIO 16
SELECT title as titulo FROM movies WHERE release_date BETWEEN "2004-01-01" AND "2008-12-31";

# EJERCICIO 17
SELECT title as titulo FROM movies WHERE rating > 3 AND awards > 1 AND release_date BETWEEN "1988-01-01" AND "2009-12-31" ORDER BY rating ASC;

# EJERCICIO 18
SELECT title as titulo FROM movies WHERE rating > 3 AND awards > 1 AND release_date BETWEEN "1988-01-01" AND "2009-12-31" ORDER BY rating DESC LIMIT 3 OFFSET 10;

select * from movies order by rating desc;
SELECT * FROM movies WHERE rating > 3 AND awards > 1 AND release_date BETWEEN "1988-01-01" AND "2009-12-31" ORDER BY rating DESC;