## a) ¿Cuál es la primary key para la tabla empleados? Justificar la respuesta.

Nº de legajo. Porque es el numero con el que está registrado en la empresa y es único.

## b) ¿Cuál es la primary key para la tabla departamentos? Justificar la respuesta.

ID departamento. De los tres campos es el único ID y debería ser único.

## c) ¿Qué relación/es existirían entre las tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de qué tabla hace referencia dicha foreign key? Justificar la respuesta.

La relacion que existe entre las tablas es la residencia de un empleado en un departamento.
En la tabla Empleado deberia haber un foreign key "id_departamento" que señale al primary key de la tabla Departamento.