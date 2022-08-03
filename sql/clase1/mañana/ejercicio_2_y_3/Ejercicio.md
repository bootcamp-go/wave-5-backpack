### Ejercicio 2
// a:
La clave primaria en empleados es el dni , ya que es unico y se puede usar como identificador.

// b:
La clave primaria en nuestra tabla departamentos es id con la propiedad auto incremental para identificar cada registro de manera única.

// c:
La relación que existiria entre ambas tablas seria poder identificar un empleado con su respectivo departamento de tal forma que en nuestra tabla empleados implementariamos un llave foranea que hace referencia a la tabla departamentos.

Esta solución nos permite realizar una sola consulta obteniendo los datos de un emplado y su departamento.
```

```sql
-- Ejemplo: agregamos datos a nuestras tablas
INSERT INTO departamentos VALUES (DEFAULT, "Recursos Humanos", "Monroe 860");
INSERT INTO departamentos VALUES (DEFAULT, "Bootcamp Go", "Monroe 860");
INSERT INTO departamentos VALUES (DEFAULT, "Investigación y Tecnología", "Monroe 860");
INSERT INTO departamentos VALUES (DEFAULT, "Global Red Team", "Monroe 860");
INSERT INTO departamentos VALUES (DEFAULT, "Administración", "Monroe 860");

INSERT INTO empleados VALUES ("50398102", "ABC123", "Marcos", "Beltran", "1990-04-01", "2022-07-01", "Backend Developer GO", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398103", "ABC124", "Franco", "Perez", "1990-04-02", "2022-07-01", "Fronted Developer ReactJs", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398104", "ABC125", "Luciana", "Rodriguez", "1990-04-03", "2022-07-01", "Backend Developer PHP", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398105", "ABC126", "Pricilla", "Sosa", "1990-04-04", "2022-07-01", "Backend Developer GO", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398106", "ABC127", "Esteban", "Blanco", "1990-04-05", "2022-07-01", "Frontend Developer JS", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398107", "ABC128", "Patricio", "Velez", "1990-04-06", "2022-07-01", "Recruiter", 90000, 80000, 1);
INSERT INTO empleados VALUES ("50398108", "ABC129", "Nahuel", "Pardo", "1990-04-07", "2022-07-01", "Analista de Seguridad", 90000, 80000, 4);
INSERT INTO empleados VALUES ("50398109", "ABC110", "Francisco", "Valle", "1990-04-08", "2022-07-01", "Investigador IT", 90000, 80000, 3);
INSERT INTO empleados VALUES ("50398110", "ABC111", "Gisel", "Pintos", "1990-04-09", "2022-07-01", "Recruiter", 90000, 80000, 1);
INSERT INTO empleados VALUES ("50398111", "ABC112", "Iam", "Charls", "1990-04-10", "2022-07-01", "Recruiter", 90000, 80000, 1);
INSERT INTO empleados VALUES ("50398112", "ABC113", "Cristian", "Valdes", "1990-04-11", "2022-07-01", "Administrador", 90000, 80000, 5);
INSERT INTO empleados VALUES ("50398113", "ABC114", "Mora", "Moreno", "1990-04-12", "2022-07-01", "Backend Developer GO", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398114", "ABC115", "Maria", "Soles", "1990-04-13", "2022-07-01", "Backend Developer GO", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398115", "ABC116", "Karim", "Benzema", "1990-04-14", "2022-07-01", "Backend Developer GO", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398116", "ABC117", "Santiago", "Lopez", "1990-04-15", "2022-07-01", "Backend Developer GO", 90000, 80000, 2);
INSERT INTO empleados VALUES ("50398117", "ABC118", "Federico", "Aimar", "1990-04-16", "2022-07-01", "Backend Developer GO", 90000, 80000, 2);

SELECT * FROM departamentos;
SELECT * FROM empleados;

-- Obteniendo datos de un empleado y su departamento mediante su dni. --
SELECT 
e.dni, concat(e.nombre, e.apellido) as nombre_apellido,
e.fecha_nac, e.fecha_ingreso, e.cargo, e.sueldo, e.neto,
d.nombre, d.direccion 
FROM empleados AS e 
INNER JOIN departamentos AS d 
ON e.departamentos_id = d.id 
WHERE e.dni = 50398117;
```

