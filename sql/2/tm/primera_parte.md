## ¿A qué se denomina JOIN en una base de datos?
Combina datos de dos tablas a partir de una o varias condiciones en comun

## Nombre y explique 2 tipos de JOIN.
Left Join: Incluye todos los datos de la primer tabla (izquierda) y solo los datos de la segunda tabla (derecha) que cumplen con las condiciones en comun
Full Join: Incluye todos los datos de ambas tablas si existe al menos una condicion en comun

## ¿Para qué se utiliza el GROUP BY?
Se utiliza para agrupar filas que tienen valores en comun. Permite realizar operaciones de agregacion con los datos agrupados

## ¿Para qué se utiliza el HAVING?
Se utiliza para filtrar los grupos

## Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno
Imagen izquierda: Inner join (join)
Imagen derecha: Left join

## Escriba una consulta genérica por cada uno de los diagramas a continuación
Imagen izquierda:
    select * from A
    right join B on A.id = B.a_id

Imagen derecha:
    select * from A
    full outer join B on A.id = B.a_id