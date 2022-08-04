USE movies_db;
SELECT se.title titulo, ge.name nombre 
FROM series se 
LEFT JOIN genres ge ON se.genre_id = ge.id;