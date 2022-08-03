	#Ej2
    SELECT * FROM movies;
    
    #Ej3
    SELECT first_name, last_name, rating FROM actors;
    
    #Ej4
    SELECT title AS titulo
    FROM series AS series;
    
    #Ej5
    SELECT first_name, last_name FROM actors
    WHERE rating > 7.5;
    
    #Ej6
    SELECT title, rating, awards FROM movies
    WHERE rating > 7.5 AND awards > 2;
    
    #Ej7
    SELECT title, rating FROM movies
    ORDER BY rating;
    
    #Ej8
    SELECT title FROM movies
    LIMIT 3;
    
    #Ej9
    SELECT title FROM movies
    ORDER BY rating DESC
    LIMIT 5;
    
    #Ej10
    SELECT * FROM movies
    ORDER BY rating DESC
    LIMIT 5 OFFSET 5;
    
    #Ej11
    SELECT * FROM actors
    LIMIT 10;
    
    #Ej12
    SELECT * FROM actors
    LIMIT 10 OFFSET 20;
    
    #Ej13
    SELECT * FROM actors
    LIMIT 10 OFFSET 40;
    
    #Ej14
    SELECT title, rating FROM movies
    WHERE title = "Toy Story";
    
    #Ej15
    SELECT * FROM actors
    WHERE first_name LIKE "Sam%";
    
    #Ej16
    SELECT title FROM movies
    WHERE release_date BETWEEN '2004-01-01 00:00:00' AND '2008-01-01 00:00:00';
    
    #Ej17
    SELECT title FROM movies
    WHERE rating > 3
    AND awards > 1
    AND release_date BETWEEN '1988-01-01 00:00:00' AND '2009-01-01 00:00:00';
    
    #Ej18
    SELECT title FROM movies
    WHERE rating > 3
    AND awards > 1
    AND release_date BETWEEN '1988-01-01 00:00:00' AND '2009-01-01 00:00:00'
    LIMIT 3 OFFSET 10;