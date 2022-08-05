-- Primera Parte

-- 1. Se denomina JOIN a la sentencia que genera una intersección, la cual puede ser de diferentes tipos y bajo ciertas condiciones, entre dos columnas de diferentes tablas.
-- 2.1 INNER JOIN: es aquella sentencia que genera la intersección entre dos columnas que comparten una condición en común.
-- 2.2 LEFT JOIN: es aquella sentencia que trae los datos de la columna izquierda y que por otra parte comparte una condición con la columna derecha. Si la columna izquierda
	-- no comparte ninguna condición con la columna derecha, entonces se devuelve el valor null.
-- 3. Se utiliza para agrupar datos de una columna, a partir de la verificación y cumplimientos de ciertas condiciones.
-- 4. Se utiliza para hacer una filtración de los datos obtenidos a partir de un Group By.alter
-- 5.1 Corresponde a un inner join.
-- 5.2 Corresponde a un left join.
-- 6.1 

SELECT act.* FROM movies_db.actors act;


SELECT mo.*, ac.first_name, act.last_name
	FROM movies_db.movies mo
    INNER JOIN actors act 
    ON mo.id=ac.favorite_movie_id;





