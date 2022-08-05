## ¿Cuantas colecciones tiene la base de datos?
hay 1 coleccion. 'restaurantes'

## ¿Cuantos elementos hay en cada coleccion? ¿Cuanto pesa cada coleccion?
hay 25359 elementos y la coleccion pesa 4304896 bytes

## ¿Cuantos indices en cada coleccion? ¿Cuanto espacio ocupan los indices de cada coleccion?
hay 1 indice. '_id'. los indices pesan 266240 bytes

## Traer un documento de ejemplo de cada coleccion
{ _id: ObjectId("5eb3d668b31de5d588f429c3"),
  direccion: 
   { edificio: '71-24',
     coord: [ -73.8221418, 40.7272376 ],
     calle: 'Main Street',
     codigo_postal: '11367' },
  barrio: 'Queens',
  tipo_cocina: 'Jewish/Kosher',
  grados: 
   [ { date: 2014-05-07T00:00:00.000Z, grado: 'A', puntaje: 9 },
     { date: 2013-09-04T00:00:00.000Z, grado: 'B', puntaje: 15 },
     { date: 2013-03-21T00:00:00.000Z, grado: 'A', puntaje: 11 },
     { date: 2012-05-23T00:00:00.000Z, grado: 'A', puntaje: 12 },
     { date: 2011-11-01T00:00:00.000Z, grado: 'A', puntaje: 11 } ],
  nombre: 'Shimons Kosher Pizza',
  restaurante_id: '40366586' }

## Para cada coleccion, listar los campos a nivel raiz (ignorar campos dentro de documentos anidados) y sus tipos de datos.

_id
direccion
barrio
tipo_cocina
grados
nombre
restaurante_id