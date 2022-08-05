## Devolver el restaurante_id, nombre, barrio y tipo_cocina pero excluyendo _id para un documento (el primero)
db.restaurantes.findOne(
    {},
    {restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0}
)

{ barrio: 'Manhattan',
  tipo_cocina: 'American',
  nombre: 'Cafe Metro',
  restaurante_id: '40363298' }

## Devolver el restaurante_id, nombre, barrio y tipo_cocina para los primeros 3 restaurantes que contengan 'Bake' en alguna parte de su nombre
db.restaurantes.find(
    {nombre: /Bake/},
    {restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0}
).limit(3)

{ barrio: 'Staten Island',
  tipo_cocina: 'American',
  nombre: 'Perkins Family Restaurant & Bakery',
  restaurante_id: '40370910' }
{ barrio: 'Queens',
  tipo_cocina: 'Caribbean',
  nombre: 'Western Bakery',
  restaurante_id: '40377560' }
{ barrio: 'Bronx',
  tipo_cocina: 'Bakery',
  nombre: 'Morris Park Bake Shop',
  restaurante_id: '30075445' }

## Contar los restaurantes de comida (tipo_cocina) china (Chinise) o tailandesa (Thai) del barrio (barrio) Bronx. Consultar or versus in
db.restaurantes.countDocuments(
  {tipo_cocina: {$in: ["Chinise", "Thai"]}, barrio: "Bronx"}
  )

2

db.restaurantes.countDocuments(
  {$or: [{tipo_cocina: "Chinise"}, {tipo_cocina: "Thai"}], barrio: "Bronx"}
  )

2

