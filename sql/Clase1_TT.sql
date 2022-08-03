/*1*/
SELECT * 
FROM movies_db.movies ;

/*2*/
SELECT first_name, last_name, rating 
FROM movies_db.actors ;

/*3*/
SELECT s.title TITULO 
FROM series s ;

/*4*/
SELECT FIRST_NAME NOMBRE, LAST_NAME APELLIDO FROM MOVIES_DB.ACTORS 
WHERE RATING > 7.5 ;

/*5*/
SELECT title, rating, awards from movies_db.movies 
WHERE RATING > 7.5
AND AWARDS > 2 ;

/*6*/
SELECT title, rating from movies_db.movies
ORDER BY rating ;

/*7*/
SELECT title FROM movies
LIMIT 3 ;

/*8*/
SELECT title, rating
FROM movies_db.movies
ORDER BY rating DESC
LIMIT 5 ;

/*9*/
SELECT title, rating
FROM movies_db.movies
ORDER BY rating DESC
LIMIT 5 
OFFSET 4 ;

/*10*/
SELECT *
FROM movies_db.actors
LIMIT 10 ;

/*11*/
SELECT *
FROM movies_db.actors
LIMIT 10 
OFFSET 29;

/*12*/
SELECT *
FROM movies_db.actors
LIMIT 10 
OFFSET 49;

/*13*/
SELECT title, rating
FROM movies_db.movies
WHERE title LIKE 'Toy Story' ;

/*14*/
SELECT *
FROM movies_db.actors
WHERE first_name LIKE 'Sam%' ;

/*15*/
SELECT title, release_date
FROM movies_db.movies
WHERE release_date BETWEEN '20040101' AND '20081231' ;

/*16*/
SELECT title
FROM movies_db.movies
WHERE rating > 3
AND awards > 1
AND release_date BETWEEN '19880101' AND '20091231' 
ORDER BY rating ;

/*17*/
SELECT title
FROM movies_db.movies
WHERE rating > 3
AND awards > 1
AND release_date BETWEEN '19880101' AND '20091231' 
ORDER BY rating 
LIMIT 3
OFFSET 9;