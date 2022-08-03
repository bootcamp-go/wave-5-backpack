# PRIMERA PARTE

~¿A qué se denomina JOIN en una base de datos?
    Al comando que nos permite unir la interseccion entre conjuntos de datos bajo ciertas condiciones en comun.
~Nombre y explique 2 tipos de JOIN.
    -Inner Join:  Nos permite agrupar las condiciones en comun que comparten 2 conjuntos.
    -Left Join: Nos permiter agrupar las condiciones en comun que tiene el conjunto IZQ con el conjunto DER.
~¿Para qué se utiliza el GROUP BY?
    -Nos permite resumir el resultado de una consulta, agrupando de acuerdo a las filas indicadas.
~¿Para qué se utiliza el HAVING?
    -Se utiliza dentro de los grupos del Group By permitiendonos filtrar la data.
~Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
    a) INNER JOIN. b) LEFT JOIN.
Escriba una consulta genérica por cada uno de los diagramas a continuación:
    a)  SELECT *
        FROM sample sp
        RIGHT JOIN test ts
        ON sp.id = ts.service_price;
    
    b)  SELECT *
        FROM sample1
        FULL JOIN sample2
        ON sample1.mock = sample2.mock
    

