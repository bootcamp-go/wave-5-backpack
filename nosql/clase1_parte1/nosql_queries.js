// EJERCICIO 1

/* 1. Devolver restaurante_id, nombre, barrio y tipo_cocina pero excluyendo _id para un documento 
(el primero). */

db.restaurantes.find(
	{},
	{ restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0 }
).limit(1)


/* 2. Devolver restaurante_id, nombre, barrio y tipo_cocina para los primeros 3 restaurantes 
que contengan 'Bake' en alguna parte de su nombre. */

db.restaurantes.find(
	{ nombre: /Bake/ },
	{ restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0 }
).limit(3)

/* 3. Contar los restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) del 
barrio (barrio) Bronx. Consultar or versus in. */

db.restaurantes.find( 
	{ tipo_cocina: { $in: ["Chinese", "Thai"] }, barrio: "Bronx" }  
).count()


// EJERCICIO 2

/* 1. Traer 3 restaurantes que hayan recibido al menos una calificación de grado 'A' con 
puntaje mayor a 20. Una misma calificación debe cumplir con ambas condiciones simultáneamente; 
investigar el operador elemMatch. */

db.restaurantes.find(
	{ grados : { $elemMatch: { grado: "A", puntaje: { $gt: 20 } } } }
).limit(3)

/* 2. ¿A cuántos documentos les faltan las coordenadas geográficas? En otras palabras, 
revisar si el tamaño de direccion.coord es 0 y contar. */

db.restaurantes.find(
	{ "direccion.coord": { $size: 0 } }
).count()

/* 3. Devolver nombre, barrio, tipo_cocina y grados para los primeros 3 restaurantes; 
de cada documento solo la última calificación. Ver el operador slice. */

db.restaurantes.find(
	{},
	{ barrio: 1, tipo_cocina: 1, grados: { $slice: -1 }, _id: 0 }
).limit(3)

