USE movies_db;

EXPLAIN DELETE FROM movies_temp
WHERE awards < 5;