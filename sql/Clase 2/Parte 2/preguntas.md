# Preguntas
## 1. ¿A qué se denomina JOIN en una base de datos?
Un Join es la union de 2 tablas
## 2. Nombre y explique 2 tipos de JOIN.
- Inner Join: Muestra solo la interseccion entre 2 tablas
- Full Join: Muestra la union de 2 tablas, tanto las tablas en si, como su interseccion
 ## 3. ¿Para qué se utiliza el GROUP BY?
 Sirve para agrupar filas que tengan el mismo valor
 ## 4. ¿Para qué se utiliza el HAVING?
Funciona de la misma forma que _WHERE_, con la diferencia de que _HAVING_ se usa para funciones agregadas
## 5. Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
1. Inner Join
2. Left Join
## 6. Escriba una consulta genérica por cada uno de los diagramas a continuación:
1.
``` sql 
SELECT * FROM TableA 
RIGHT JOIN TableB
WHERE TableA.B_id = TableB.ID
```
2. 
``` sql
SELECT * FROM TableA
LEFT JOIN TableB
WHERE TableA.B_id = TableB.ID
UNION 
SELECT * FROM TableA 
RIGHT JOIN TableB 
WHERE TableA.B_id = TableB.ID
```