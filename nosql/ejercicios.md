# Para Empezar

1. ¿Cuántas colecciones tiene la base de datos?

Tiene 1

2. ¿Cuántos documentos en cada colección? ¿Cuánto pesa cada colección?

```javascript
db.restaurantes.count()
```

3. ¿Cuántos índices en cada colección? ¿Cuánto espacio ocupan los índices de cada colección?

1 Indice que pesa 260 KB

4. Traer un documento de ejemplo de cada colección. db.collection.find(...).pretty() nos da un formato más legible.

Obtenemos el primer documento de la coleccion utilizando la query: 

```javascript
db.restaurantes.find().limit(1).pretty()
```

5. Para cada colección, listar los campos a nivel raíz (ignorar campos dentro de documentos anidados) y sus tipos de datos.

```javascript
'_id: object'
'direccion: object'
'barrio: string'
'tipo_cocina: string'
'grados: object'
'nombre: string'
'restaurante_id: string'
```

# Ejercicio 1: NoSQL


1. Devolver restaurante_id, nombre, barrio y tipo_cocina pero excluyendo _id para un documento (el primero).

```javascript
db.restaurantes.find({}, {restaurante_id:1, nombre:1, barrio:1, tipo_cocina:1, _id:0})
```

2. Devolver restaurante_id, nombre, barrio y tipo_cocina para los primeros 3 restaurantes que contengan 'Bake' en alguna parte de su nombre.

```javascript
db.restaurantes.find({nombre:/Bake/}, {restaurante_id:1, nombre:1, barrio:1, tipo_cocina:1, _id:0}).limit(3)
```

3. Contar los restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) del barrio (barrio) Bronx. Consultar or versus in.

```javascript
db.restaurantes.count({barrio:"Bronx",tipo_cocina:{$in:["Chinese","Thai"]}})
```

# Ejercicio 2: NoSQL

1. Traer 3 restaurantes que hayan recibido al menos una calificación de grado 'A' con puntaje mayor a 20. Una misma calificación debe cumplir con ambas condiciones simultáneamente; investigar el operador elemMatch.

```javascript
db.restaurantes.find({grados: {$elemMatch: {grado:"A", puntaje: {$gt: 20}}}}).limit(3)
```

2. ¿A cuántos documentos les faltan las coordenadas geográficas? En otras palabras, revisar si el tamaño de direccion.coord es 0 y contar.

```javascript
db.restaurantes.count({"direccion.coord" : {$size: 0}})
```

3. Devolver nombre, barrio, tipo_cocina y grados para los primeros 3 restaurantes; de cada documento solo la última calificación. Ver el operador slice.

```javascript
db.restaurantes.find({},{nombre:1, barrio:1, tipo_cocina:1, grados: {$slice: -1}}).limit(3)
```

# Ejercicio 1: Popurrí


1. ¿Cuál es top 3 de tipos de cocina (cuisine) que podemos encontrar entre los datos? Googlear "mongodb group by field, count it and sort it". Ver etapa limit del pipeline de agregación.

```javascript
db.restaurantes.aggregate([
  {
     $group: { 
        _id: "$tipo_cocina", 
        countA: { $sum: 1}
     }
  },
  {
     $sort:{countA:-1}
  },{ $limit : 3 }
])
```

2. Cuáles son los barrios más desarrollados gastronómicamente? Calcular el promedio ($avg) de puntaje (grades.score) por barrio; considerando restaurantes que tengan más de tres reseñas; ordenar barrios con mejor puntaje arriba. Ayuda:
  a. match es una etapa que filtra documentos según una condición, similar a db.orders.find(<condición>).
  b. Parece necesario deconstruir las listas grades para producir un documento por cada puntaje utilizando la
  etapa unwind.

```javascript
db.restaurantes.aggregate([
  {
     $unwind: {path: "$grados"}
  },
  {
   $match : {'grados.puntaje' :{ $gt : 3 } }
  },
  {
    $group:{
      _id: "$barrio",
      avgGrades: { $avg: "$grados.puntaje" }

    }
  },
  {
    $sort:{avgGrades:-1}
  }
])
```

3. Una persona con ganas de comer está en longitud -73.93414657 y latitud 40.82302903, ¿qué opciones tiene en 500 metros a la redonda? Consultar geospatial tutorial.

en centerSphere: [ [x,y], radio(en millas) / 3963.2]

```javascript
db.restaurantes.find(
  { "direccion.coord":
   { $geoWithin:
      { $centerSphere: [ [ -73.93414657, 40.82302903 ], 0.3107 / 3963.2 ] 
    } 
  } 
},
{restaurante_id:1, nombre:1, barrio:1, tipo_cocina:1, _id:1}
)

```
