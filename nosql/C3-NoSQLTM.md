# Práctica MongoDB

## Práctica

### ¿Cuántas colecciones tiene la base de datos?

```sql
db.stats()
```

> 1 Collection

### ¿Cuántos documentos hay en cada colección?

```sql
db.restaurantes.count()
```

> Tiene **25359** documentos

### ¿Cuánto pesa cada colección?

```sql
db.restaurantes.dataSize()
```

> **11140976** bytes

### ¿Cuántos índices en cada colección?

> 1 index (\_id)

### ¿Cuánto espacio ocupan los índices de cada colección?

> 266240 bytes

### Traer un documento de ejemplo de cada colección

```sql
db.restaurantes.find().pretty()
```

### Para cada colección, listar los campos a nivel raíz (ignorar campos dentro de documentos anidados) y sus tipos de datos

---

## Ejercicio 1: SQL

1. Devolver restaurante_id, nombre, barrio y tipo_cocina pero excluyendo \_id para un documento (el primero).

   ```sql
   db.restaurantes.findOne({}, restaurante_id:1, nombre:1, barrio: 1, tipo_cocina:1, _id_0)
   db.getCollection('restaurantes').findOne({},{ restaurante_id : 1, nombre : 1, barrio : 1,  tipo_cocina : 1,  _id : 0 })
   ```

2. Devolver restaurante_id, nombre, barrio y tipo_cocina para los primeros 3 restaurantes que contengan 'Bake' en alguna parte de su nombre.

   ```sql
   db.restaurantes.find({nombre : /Bake/},{ restaurante_id : 1, nombre: 1, barrio: 1,  tipo_cocina : 1,  _id : 0 }).limit(3)
   db.getCollection('restaurantes').find({nombre : /Bake/},{ restaurante_id : 1, nombre: 1, barrio: 1,  cuisine : 1,  _id : 0 }).limit(3)
   ```

3. Contar los restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) del barrio (barrio) Bronx.

   ```sql
   db.restaurantes.count({tipo_cocina : {$in : ["Chinese", "Thai"]}, barrio: "Bronx"})
   db.getCollection('restaurantes').count({tipo_cocina : {$in : ["Chinese", "Thai"]}, barrio: "Bronx"})
   ```

---

## Ejercicio 2: NoSQL

1. Traer 3 restaurantes que hayan recibido al menos una calificación de grado 'A' con puntaje mayor a 20. Una misma calificación debe cumplir con ambas condiciones simultáneamente

   ```sql
   db.restaurantes.find({ grados: { $elemMatch: { grado : "A", puntaje : { $gt : 20 } } } } ).limit(3)
   db.getCollection('restaurantes').find({ grados: { $elemMatch: { grado : "A", puntaje : { $gt : 20 } } } } ).limit(3)
   ```

2. ¿A cuántos documentos les faltan las coordenadas geográficas? En otras palabras, revisar si el tamaño de direccion.coord es 0 y contar.

   ```sql
   db.restaurantes.count( { "direccion.coord" : { $size : 0 } } )
   ```

3. Devolver nombre, barrio, tipo_cocina y grados para los primeros 3 restaurantes; de cada documento solo la última calificación.

   ```sql
   db.restaurantes.find({}, {nombre:1, barrio:1, tipo_cocina:1, grados :{$slice:-1}, _id:0}).limit(3)

   ```

## Ejercicio 3: Popurri

1. ¿Cuál es top 3 de tipos de cocina (cuisine) que podemos encontrar entre los datos? Googlear "mongodb group by field, count it and sort it".

   ```sql
   db.restaurantes.aggregate([{"$group" : {_id: {tipo_cocina: "$tipo_cocina"}, count: {$sum: 1}}},{$sort: {"count": -1}},{$limit: 3}])
   ```

2. ¿Cuáles son los barrios más desarrollados gastronómicamente? Calcular el promedio ($avg) de puntaje (grades.score) por barrio; considerando restaurantes que tengan más de tres reseñas; ordenar barrios con mejor puntaje arriba.

   ```sql
   db.restaurantes.aggregate([
   {
      $unwind: {path: "$grados"}
   },
   {
      $match : {'grados.puntaje' :{ $gt : 3 } }
   },
   {
      $group: {
         _id : "$barrio",
         promedio : {$avg : "$grados.puntaje"}
      }
   },
   {
      $sort:{
         promedio: -1
      }
   }
   ])

   ```

3. Una persona con ganas de comer está en longitud -73.93414657 y latitud 40.82302903, ¿qué opciones tiene en 500 metros a la redonda?



