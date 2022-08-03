USE movies_db;
SELECT title titulo, release_date lanzamiento FROM movies WHERE release_date BETWEEN '2004-07-04 00:00:00' AND '2008-07-04 00:00:00' ORDER BY release_date;