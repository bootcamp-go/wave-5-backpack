### ¿Cuántas colecciones tiene la base de datos?

```sql
db.stats()
```

> 1 Collection

### ¿Cuántos documentos hay en cada colección?

```sql
db.restaurantes.count()
```

>Tiene **25359** documentos

### ¿Cuánto pesa cada colección?

```sql
db.restaurantes.dataSize()
```

>11140976 bytes

### ¿Cuántos índices en cada colección?

> 1 index (_id)

### ¿Cuánto espacio ocupan los índices de cada colección?

> 266240 bytes

### Traer un documento de ejemplo de cada colección

```sql
db.restaurantes.find().pretty()
```

### Para cada colección, listar los campos a nivel raíz (ignorar campos dentro de documentos anidados) y sus tipos de datos.


#Ejercicio 1: SQL

1. Devolver restaurante_id, nombre, barrio y tipo_cocina pero excluyendo _id para un documento (el primero).

```
db.restaurantes.findOne({}, restaurante_id:1, nombre:1, barrio: 1, tipo_cocina:1, _id_0})
```

Otra opción
```
db.getCollection('restaurantes').findOne({},{ restaurant_id : 1, nombre : 1, barrio : 1,  tipo_cocina : 1,  _id : 0 })
```