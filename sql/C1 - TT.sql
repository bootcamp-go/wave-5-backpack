/* 2 */  SELECT * FROM movies;

/* 3 */  SELECT first_name nombre, last_name apellido, rating 
		 FROM actors; 
         
/* 4 */  SELECT title titulo FROM series; 

/* 5 */  SELECT first_name nombre, last_name apellido 
		 FROM actors 
         WHERE rating > 7.5; 

/* 6 */  SELECT title titulo, rating, awards premios 
		 FROM movies 
         WHERE rating > 7.5 AND awards > 2; 
         
/* 7 */  SELECT title titulo, rating 
		 FROM movies 
         ORDER BY rating;
         
/* 8 */  SELECT title titulo 
		 FROM movies 
         LIMIT 3;
         
/* 9 */  SELECT title titulo, rating 
		 FROM movies 
         ORDER BY rating DESC LIMIT 5;
         
/* 10 */ SELECT title titulo, rating 
		 FROM movies 
         ORDER BY rating DESC LIMIT 5 
         OFFSET 6;
         
/* 11 */ SELECT * FROM actors 
		 LIMIT 10;

/* 12 */ SELECT * FROM actors 
		 LIMIT 10 
         OFFSET 21;
         
/* 13 */ SELECT * FROM actors 
		 LIMIT 10 
         OFFSET 41;
         
/* 14 */ SELECT title titulo, rating 
		 FROM movies 
         WHERE title = "Toy Story";
         
/* 15 */ SELECT * FROM actors 
		 WHERE first_name LIKE "Sam%";
         
/* 16 */ SELECT title titulo, release_date 
		 FROM movies 
         WHERE release_date BETWEEN '20040101' AND '20081231';
         
/* 17 */ SELECT title titulo FROM movies 
		 WHERE rating > 3
         AND awards > 1
         AND release_date BETWEEN '19980101' AND '20091231'
         ORDER BY rating;
		
/* 18 */ SELECT title titulo FROM movies 
		 WHERE rating > 3
         AND awards > 1
         AND release_date BETWEEN '19980101' AND '20091231'
         ORDER BY rating
         LIMIT 3
         OFFSET 10;
         