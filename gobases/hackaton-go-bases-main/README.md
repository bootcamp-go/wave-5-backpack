# Bootcamp Go 
## **Hackathon Go Bases💥**
Práctica individual </br>
### <b>Objetivo</b> </br>
El objetivo de esta guía práctica es que podamos afianzar y profundizar los conceptos vistos en Go Bases. Para esto, vamos a plantear un desafío integrador que engloba los temas estudiados.  </br>
Tendrán tiempo para realizarlo hasta las 16 hs AR/UR - 15 hs CH - 14 hs CO/ME. </br>

### **Planteo** </br>
Una aerolínea pequeña necesita un sistema de reservas de pasajes a diferentes países, y requiere un archivo con la información de los pasajes sacados en las últimas 24 horas. 
Deben crear un sistema de reservas para generar esos archivos
El archivo en cuestión es del tipo valores separados por coma (csv por su siglas en inglés), donde los campos están compuestos por: id, nombre, email, país de destino, hora del vuelo y precio. 

### **¿Are you ready?** </br>


# Desafío
Realizar un programa que sirva como herramienta para calcular diferentes datos estadísticos. Para lograrlo, debes clonar este repositorio que contiene un archivo .csv con datos generados.


### ***¡Atención!*** **Los ejemplos a continuación son sólo de guía, el desafío se puede resolver de múltiples maneras.**





**Requerimiento 1 - Cargar archivos:** </br>
Implementar un módulo para leer el archivo donde se encuentran los tickets del día.</br></br>
**Requerimiento 2 - Crear:** </br>
Un método para crear un nuevo ticket añadir al registro.
```go 
func (b *bookings) Create(t Ticket) (Ticket, error) {}
``` 

**Requerimiento 3 - Leer:**</br>
Un método para traer un ticket a través de su campo id.
```go
func (b *bookings) Read(id int) (Ticket, error) {} 
```


**Requerimiento 4 - Update:**</br>
Un método para actualizar los campos de un ticket.
```go
func (b *bookings) Update(id int, t Ticket) (Ticket, error) {}
```

**Requerimiento 5 - Delete:**</br>
Un método para eliminar un campo por su id.
```go
func (b *bookings) Delete(id int) (int, error) {}
```


