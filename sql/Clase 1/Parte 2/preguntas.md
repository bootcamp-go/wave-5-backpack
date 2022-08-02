# Preguntas
## a. ¿Cuál es la primary key para la tabla _empleados_? Justificar la respuesta.
##
Numero de legajo, ya que esta es unica por cada empleado
## b. ¿Cuál es la primary key para la tabla _departamentos_? Justificar la respuesta.
ID departamento, ya que cada departamento tiene un ID asignado unico e irrepetible
## c. ¿Qué relación/es existirían entre las tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de qué tabla hace referencia dicha foreign key? Justificar la respuesta.

- Entre las tablas existe una relacion uno a muchos, donde un departamento puede tener muchos empleados
- La llave foranea debe ir en _Exmpleado_, ya que la relacion *uno-muchos*, se necesita que la entidad *muchos* (en este caso _empleado_) guarde un identificador de su entidad *uno* (_departamento_)
- AL campo ID de la tabla _Departamento_, ya que la llave foranea debe ser capaz de diferenciar entre diferentes entidades del mismo tipo, y dado que _ID Departamento_ es un identificador unico, ademas de *primary key*, el el ideal
