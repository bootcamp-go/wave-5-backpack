# ¿Cuáles son las diferencias entre White Box y Black Box? 
La White Box trata sobre hacer testing conociendo la estructura del código (unitario y integracion) y la Black Box trata sobre solamente probar la funcionalidad (requerimientos funcionales y no funcionales) sin conocer como se esta comportando por dentro el sistema.

Ejemplo de white box:
Test que invoca una función del sistema modificando sus dependencias en el codigo como una Base de datos de prueba.

Ejemplo de black box:
Probar un punto api y comprobar que su respuesta sea la que se definio en los requerimientos funcionales.

# ¿Qué es un test funcional?
Un test funcional busca validar que el requerimiento este cumpliendose como se determinó, sin la necesidad de conocer como esta comportandose el software en su interior.

# ¿Qué es un Test de Integración?
Un test que busca emular un ambiente cercano al de producción, ya que se prueban las funcionalidades que se estan comunicando con servicios externos como bases de datos, apis, etc, y se evalua que la implementación con este servicio sea exitosa.

# Indicar las dimensiones de calidad prioritarias en MELI.
Seguridad, Mantenibilidad, Funcionalidad, Rendimiento, Fiabilidad.
