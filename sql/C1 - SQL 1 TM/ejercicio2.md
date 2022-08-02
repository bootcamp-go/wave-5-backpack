# Ejercicio 2

## a. ¿Cuál es la primary key para la tabla empleados? Justificar la respuesta.
Puede ser el documento de identidad, pero para efectos del negocio, es legajo, porque es un identificador único como trabajador y representa mayor importancia a nivel empresa.


## b. ¿Cuál es la primary key para la tabla departamentos? Justificar la respuesta.
Id del departamento, porque entre los atributos es el que identifica de manera única al departamento. 


## c. ¿Qué relación/es existirían entre las tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de qué tabla hace referencia dicha foreign key? Justificar la respuesta.
Una relación de uno a muchos, debido a que en un departamento trabajan muchos empleados. La tabla que tiene la foreign key es la correspondiente a Empleados, esto, porque un departamento está vinculado a muchos empleados, y un empleado se referencia a un sólo departamento, entonces, se requiere conocer código para identificarlo.
