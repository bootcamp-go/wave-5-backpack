# Escenario
La empresa Super Software SRL desea una base de datos para gestionar los departamentos en los cuales se encuentra cada uno de sus empleados. Para ello, brinda la siguiente información en donde detalla los datos que necesita:
- De sus Empleados necesita almacenar: N.º de legajo, dni, apellido, nombre, fecha de nacimiento, fecha de incorporación, cargo, sueldo neto.
- De sus Departamentos necesita almacenar: ID departamento, nombre departamento, dirección departamento.

## Ejercicio 1
- A partir del planteo de los requerimientos de la empresa, se solicita modelar los mismos mediante un DER (Diagrama entidad relación), teniendo en cuenta las posibles entidades y relaciones necesarias. Además, se solicita crear una nueva base de datos llamada “emple_dep”. <br/>
![DER](https://user-images.githubusercontent.com/107702332/182619770-2d0459cb-b770-4be0-bd2d-57f03e800335.jpeg)


## Ejercicio 2
Una vez modelada y planteada la base de datos, responder a las siguientes preguntas:
1. ¿Cuál es la primary key para la tabla **empleados**? Justificar la respuesta.
- La PK es el dni, ya que es el número identificador unico de cada persona.
2. ¿Cuál es la primary key para la tabla **departamentos**? Justificar la respuesta.
- LA PK sería el id del departamento

3. ¿Qué relación/es existirían entre las tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de qué tabla hace referencia dicha foreign key? Justificar la respuesta.

- Para el lado del empleado la cardinalidad seria muchos (M) y del lado del departamento uno (1)
- La FK debería estar en la tabla empleados y debe hacer referencia a la PK de departamentos que sería el ID
