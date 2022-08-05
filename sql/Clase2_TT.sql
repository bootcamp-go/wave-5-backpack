SELECT * FROM movies_db.movies;

/*2*/
INSERT INTO movies_db.movies VALUES (22,NULL,NULL,"Something",2.5,0,"2020-09-04 00:00:00",110,3);

/*3*/
INSERT INTO movies_db.genres VALUES (13,"2020-09-04 00:00:00",NULL,"Peli",13,1);

/*4*/
UPDATE movies_db.movies SET genre_id = 13 where id = 22;

/*5*/
UPDATE movies_db.actors SET favorite_movie_id = 22 where id = 49;

/*6*/
CREATE TEMPORARY TABLE temp_movie (SELECT * FROM movies_db.movies);

/*7*/
DELETE from movies_db.temp_movie WHERE awards < 5;

/*8*/
SELECT g.name
FROM movies_db.genres g 
JOIN movies_db.movies m ON g.id = m.genre_id
GROUP BY m.genre_id
HAVING COUNT(m.genre_id) >= 1;

/*9*/
SELECT a.first_name, a.last_name
FROM movies_db.actors a 
JOIN movies_db.movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3;

/*12*/
CREATE INDEX movies_title_idx ON movies_db.movies (title);