## 1. ¿Cuál es la primary key para la tabla empleados? Justificar la respuesta.
- DNI, ya que es un identificador único de cada persona como su Tarjeta de Identidad.
## 2. ¿Cuál es la primary key para la tabla departamentos? Justificar la respuesta.
- ID, ya que es un identificador único de cada departamento entre las tres variables
## 3. ¿Qué relación/es existirían entre las tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de qué tabla hace referencia dicha foreign key? Justificar la respuesta.
- Aunque no vea una relación directa se puede ver con el cargo el tipo de departamento al que pertenece.
- Debería haber un foreign key en la tabla empleados para que se pueda acceder a la tabla departamentos.
- Cree un nuevo campo llamado Departamento_ID que funciona como foreign key en la tabla empleados haciendo referencia al ID de la tabla departamentos.