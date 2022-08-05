## Traer 3 restaurantes que hayan recibido al menos una calificacion de grado 'A' con puntaje mayor a 20.
## Una misma calificacion debe cumplir con ambas condiciones simultaneamente; investigar el operador elemMatch
db.restaurantes.find(
    {grados: {$elemMatch: {grado: "A", puntaje: {$gt: 20}}}}
).limit(3)

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

## ¿A cuantos documentos les faltan las coordenadas geograficas? En otras palabras, revisar si el tamaño de
## direccion.coord es 0 y contar
db.restaurantes.countDocuments(
    {"direccion.coord": {$size: 0}}
)

2


## Devolver nombre, barrio, tipo_cocina y grados para los primeros 3 restaurantes; de cada documento solo
## la ultima calificacion. Ver el operador slice
db.restaurantes.find(
    {},
    {nombre: 1, barrio: 1, tipo_cocina: 1, grados: {$slice: -1}}
).limit(3)

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