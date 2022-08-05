##  :closed_lock_with_key: Práctica MongoDB :school_satchel::books:

### Para empezar

#### ¿Cuántas colecciones tiene la base de datos?

```mysql
db.stats().collections
Colecciones: 1

db.getCollectionNames()
[ 'restaurantes' ]
```

#### ¿Cuántos documentos hay en cada colección? ¿Cuánto pesa cada colección?

```mysql
db.restaurantes.countDocuments()
Documentos en cada colección: 25359

db.restaurantes.dataSize()
Peso de cada colección: 11140976
```

#### ¿Cuántos índices en cada colección? ¿Cuánto espacio ocupan los índices de cada colección?

```mysql
db.stats().indexes
Indices: 1

db.restaurantes.getIndexes()
[ { v: 2, key: { _id: 1 }, name: '_id_' } ]

db.stats().indexSize
425984
```

#### Traer un documento de ejemplo de cada colección. db.collection.find(...).pretty() nos da un formato más legible.

```mysql
db.restaurantes.find().pretty()
```

#### Para cada colección, listar los campos a nivel raíz (ignorar campos dentro de documentos anidados) y sus tipos de datos.

```mysql
[db.restaurantes.findOne()].forEach( function(my_doc) { for (var key in my_doc) { print(key + ': ' + typeof my_doc[key]) } } )

[
  '_id',
  'direccion',
  'barrio',
  'tipo_cocina',
  'grados',
  'nombre',
  'restaurante_id'
]
```

## Ejercicio 1: SQL

#### Devolver restaurante_id, nombre, barrio y tipo_cocina pero excluyendo _id para un documento (el primero).

```mysql
db.restaurantes.findOne(
   { },
   {restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0 }
)

-- Información obtenida primer documento:
{ barrio: 'Manhattan',
  tipo_cocina: 'American',
  nombre: 'Cafe Metro',
  restaurante_id: '40363298' }
```

#### Devolver restaurante_id, nombre, barrio y tipo_cocina para los primeros 3 restaurantes que contengan 'Bake' en alguna parte de su nombre.

```mysql
db.restaurantes.find(
   { nombre: /Bake/ },
   { restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1 }
).limit(3)

-- Información obtenida de primeros tres restaurantes:
{ _id: ObjectId("5eb3d668b31de5d588f42a67"),
  barrio: 'Staten Island',
  tipo_cocina: 'American',
  nombre: 'Perkins Family Restaurant & Bakery',
  restaurante_id: '40370910' }
{ _id: ObjectId("5eb3d668b31de5d588f42aea"),
  barrio: 'Queens',
  tipo_cocina: 'Caribbean',
  nombre: 'Western Bakery',
  restaurante_id: '40377560' }
{ _id: ObjectId("5eb3d668b31de5d588f4292e"),
  barrio: 'Bronx',
  tipo_cocina: 'Bakery',
  nombre: 'Morris Park Bake Shop',
  restaurante_id: '30075445' }
```

#### Contar los restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) del barrio (barrio) Bronx. Consultar or versus in.

```mysql
db.restaurantes.count(
   { tipo_cocina: { $in: [ "Chinese", "Thai" ] }, barrio: "Bronx" }
)

-- Restaurantes de comida china o tailandesa del barrio Bronx:
325
```

## Ejercicio 2: NoSQL

#### Traer 3 restaurantes que hayan recibido al menos una calificación de grado 'A' con puntaje mayor a 20. Una misma calificación debe cumplir con ambas condiciones simultáneamente; investigar el operador elemMatch.

```mysql
db.restaurantes.find(
   { grados: { $elemMatch: { grado:"A", puntaje: { $gt: 20 } } }}
).limit(3)

-- 3 restaurantes que han recibido al menos una calificación de grado 'A' con puntaje mayor a 20:

{ _id: ObjectId("5eb3d668b31de5d588f42eec"),
  direccion: 
   { edificio: '107-23',
     coord: [ -73.834012, 40.683833 ],
     calle: 'Liberty Avenue',
     codigo_postal: '11417' },
  barrio: 'Queens',
  tipo_cocina: 'Caribbean',
  grados: 
   [ { date: 2014-03-29T00:00:00.000Z, grado: 'A', puntaje: 27 },
     { date: 2013-06-12T00:00:00.000Z, grado: 'A', puntaje: 12 },
     { date: 2012-05-10T00:00:00.000Z, grado: 'A', puntaje: 13 },
     { date: 2011-12-29T00:00:00.000Z, grado: 'A', puntaje: 13 } ],
  nombre: 'Gemini\'S Lounge',
  restaurante_id: '40511696' }
{ _id: ObjectId("5eb3d668b31de5d588f42fb1"),
  direccion: 
   { edificio: '892',
     coord: [ -73.92434209999999, 40.8281502 ],
     calle: 'Gerard Avenue',
     codigo_postal: '10452' },
  barrio: 'Bronx',
  tipo_cocina: 'Caribbean',
  grados: 
   [ { date: 2014-10-15T00:00:00.000Z, grado: 'A', puntaje: 21 },
     { date: 2014-05-09T00:00:00.000Z, grado: 'A', puntaje: 4 },
     { date: 2013-10-28T00:00:00.000Z, grado: 'A', puntaje: 13 },
     { date: 2013-05-29T00:00:00.000Z, grado: 'A', puntaje: 13 },
     { date: 2012-04-30T00:00:00.000Z, grado: 'A', puntaje: 11 },
     { date: 2011-12-13T00:00:00.000Z, grado: 'B', puntaje: 15 } ],
  nombre: 'Feeding Tree Style West Indian Restaurant',
  restaurante_id: '40541088' }
{ _id: ObjectId("5eb3d668b31de5d588f43f43"),
  direccion: 
   { edificio: '145',
     coord: [ -73.9805339, 40.7629624 ],
     calle: 'West   53 Street',
     codigo_postal: '10019' },
  barrio: 'Manhattan',
  tipo_cocina: 'Italian',
  grados: 
   [ { date: 2014-06-18T00:00:00.000Z, grado: 'A', puntaje: 13 },
     { date: 2013-06-24T00:00:00.000Z, grado: 'A', puntaje: 10 },
     { date: 2013-01-16T00:00:00.000Z, grado: 'B', puntaje: 21 },
     { date: 2012-07-03T00:00:00.000Z, grado: 'B', puntaje: 21 },
     { date: 2012-01-11T00:00:00.000Z, grado: 'A', puntaje: 24 } ],
  nombre: 'Remi',
  restaurante_id: '41118090' }
```

#### ¿A cuántos documentos les faltan las coordenadas geográficas? En otras palabras, revisar si el tamaño de direccion.coord es 0 y contar.

```mysql
db.restaurantes.count(
   { "direccion.coord": { $size: 0 } }
)

-- Documentos a los que les falta las coordenadas geográficas:
2
```

#### Devolver nombre, barrio, tipo_cocina y grados para los primeros 3 restaurantes; de cada documento solo la última calificación. Ver el operador slice.

```mysql
db.restaurantes.find(
   { },
   { nombre: 1, barrio: 1, tipo_cocina: 1, grados: { $slice: -1 } }
).limit(3)

-- Infomación de 3 primeros restaurantes (sólo la última calificación de cada documento):

{ _id: ObjectId("5eb3d668b31de5d588f4294f"),
  barrio: 'Manhattan',
  tipo_cocina: 'American',
  grados: [ { date: 2011-09-09T00:00:00.000Z, grado: 'A', puntaje: 13 } ],
  nombre: 'Cafe Metro' }
{ _id: ObjectId("5eb3d668b31de5d588f42930"),
  barrio: 'Queens',
  tipo_cocina: 'American',
  grados: [ { date: 2012-02-10T00:00:00.000Z, grado: 'A', puntaje: 13 } ],
  nombre: 'Brunos On The Boulevard' }
{ _id: ObjectId("5eb3d668b31de5d588f42955"),
  barrio: 'Manhattan',
  tipo_cocina: 'Pizza',
  grados: [ { date: 2011-09-26T00:00:00.000Z, grado: 'A', puntaje: 0 } ],
  nombre: 'Domino\'S Pizza' }
```

## Ejercicio 3: Popurri

#### ¿Cuál es top 3 de tipos de cocina (cuisine) que podemos encontrar entre los datos? Googlear "mongodb group by field, count it and sort it". Ver etapa limit del pipeline de agregación.

```mysql
db.restaurantes.aggregate([
  {
     $group: { 
        _id: "$tipo_cocina"
     }
  },
  {
    $sort:{'grado.puntaje':-1}
  },
  {
    $limit: 3
  }
])
```

#### ¿Cuáles son los barrios más desarrollados gastronómicamente? Calcular el promedio ($avg) de puntaje (grades.score) por barrio; considerando restaurantes que tengan más de tres reseñas; ordenar barrios con mejor puntaje arriba.

```mysql
db.restaurantes.aggregate([
  {
    $unwind: {
      path: "$grados"
    }
  },
  {
    $group: {
      _id: {
        date: "$barrio",
      },
      avg: {
        $avg: "$grados.puntaje"
      }
    }
    $sort:{'avg':-1}
  }
])
```

#### Una persona con ganas de comer está en longitud -73.93414657 y latitud 40.82302903, ¿qué opciones tiene en 500 metros a la redonda? Consultar geospatial tutorial.

```mysql
db.restaurantes.find(
  { "direccion.coord":
   { $geoWithin:
      { $centerSphere: [ [ -73.93414657, 40.82302903 ], 0.3107 / 3963.2 ] 
    } 
  } 
})
```