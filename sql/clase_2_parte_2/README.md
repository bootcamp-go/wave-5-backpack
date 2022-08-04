¿A qué se denomina JOIN en una base de datos?
Un JOIN es la intersección de dos grupos de datos por una columna en común dando como resultado los datos de ambos grupos.
Nombre y explique 2 tipos de JOIN.

INNER JOIN solo muestra los datos que tienen coincidencia en el valor de la columna en la que los grupos fueron interseccionados.
LEFT JOIN muestra todos los valores del grupo izquierdo además de mostrar los datos que tienen coincidencia en el valor de la columna en la que ambos grupos fueron interseccionados.

¿Para qué se utiliza el GROUP BY?
Dentro de una tabla agrupa las filas con el mismo valor en la columna elegida en la sentencia del GROUP BY, esto con el objetivo de obtener una agregación dividada en conjuntos de datos.

¿Para qué se utiliza el HAVING?
Sirve para hacer filtros cuando se utiliza una segregación similar al comportamiento de la sentencia WHERE.

Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
1. SELECT table_a.name, table_a.table_b FROM table_a INNER JOIN table_b ON table_a.table_b = table_b.id
2. SELECT table_a.name, table_a.table_b FROM table_a LEFT JOIN table_b ON table_a.table_b = table_b.id

Escriba una consulta genérica por cada uno de los diagramas a continuación:
1. SELECT table_a.name, table_a.table_b FROM table_a RIGHT JOIN table_b ON table_a.table_b = table_b.id
2. SELECT table_a.name, table_a.table_b FROM table_a LEFT JOIN table_b ON table_a.table_b = table_b.id UNION SELECT table_a.name, table_a.table_b FROM table_a RIGHT JOIN table_b ON table_a.table_b = table_b.id