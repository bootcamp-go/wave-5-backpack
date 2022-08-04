USE movies_db;

SELECT COUNT(*), s.title titulo
FROM series s
INNER JOIN seasons se ON se.serie_id = s.id
GROUP BY s.title;