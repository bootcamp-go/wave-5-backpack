# SQL 3 - MOVIES DB

Practica grupal

1. Explicar el concepto de normalización y para que se utiliza.

> Básicamente, la normalización está encaminada a eliminar redundancias e inconsistencias de dependencias en el diseño de las tablas.

2. Agregar una película a la tabla movies.

```sql
INSERT INTO movies VALUES (DEFAULT, null, null, "Mi peli", 10.0, 4, "2022-07-20 00:00:00", 120, 1);
```

3. Agregar un género a la tabla genres.

```sql
INSERT INTO genres VALUES(DEFAULT, null, null, "mi genero unico", 13, 1);
```

4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.

```sql
UPDATE movies
SET genre_id = 13
WHERE (title = "Mi peli x2" AND rating = 10.0);
```

5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.

```sql
UPDATE actors
SET favorite_movie_id = 26
WHERE (first_name = "Sam" AND last_name = "Worthington");
```

6. Crear una tabla temporal copia de la tabla movies.

7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

8. Obtener la lista de todos los géneros que tengan al menos una película.

9. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 

10. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

11. ¿Qué son los índices? ¿Para qué sirven?

12. Crear un índice sobre el nombre en la tabla movies.

13. Chequee que el índice fue creado correctamente.