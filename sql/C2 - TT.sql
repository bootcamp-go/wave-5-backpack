/* 1 */
/* 
La normalización es un proceso de estandarización y validación de datos que 
consiste en eliminar las redundancias o inconsistencias, completando datos mediante
una serie de reglas que actualizan la información, protegiendo su integridad y favoreciendo
la interpretación, para que así sea más simple de consultar y más eficiente para quien
la gestiona. 
*/

/* 2 */ 
INSERT INTO movies (title, rating, awards, release_date)
VALUES ("La dama y el Vagabundo", 7.7, 3, "20220311");
/* 3 */
INSERT INTO genres (name, ranking, active)
VALUES ("Romantico", 13, true);

/* 4 */
UPDATE movies
SET genre_id = (SELECT id FROM genres WHERE ranking = 13)
WHERE movies.title = "La dama y el Vagabundo";

/* 5 */
UPDATE actors,
(SELECT id FROM actors LIMIT 1) AS actor
SET favorite_movie_id = (SELECT id FROM movies WHERE title = "La dama y el Vagabundo")
WHERE actors.id = actor.id;

/* 6 */
CREATE TEMPORARY TABLE movies_cpy
SELECT * FROM movies;

/* 7 */
DELETE FROM movies_cpy
WHERE awards < 5;

/* 8 */
SELECT ge.* FROM genres ge
	JOIN movies mo ON mo.genre_id = ge.id
GROUP BY ge.id
HAVING count(mo.title) >= 1;

/* 9 */
SELECT ac.* FROM actors ac
	JOIN movies mo ON ac.favorite_movie_id = mo.id
WHERE mo.awards > 3;

/* 11 */
/*
Son un mecanismo para optimizar consultas en SQL.
Mejoran sustancialmente los tiempos de respuesta en Queries complejas.
Mejoran el acceso a los datos al proporcionar una ruta más directa a los registros. 
Evitan realizar escaneos (barridas) completas o lineales de los datos en una tabla.
*/

/* 12 */
CREATE INDEX title_idx
ON movies (title);

/* 13 */
SHOW INDEX FROM movies





