# Preguntas
## 1. ¿Cuántas colecciones tiene la base de datos?
Solo tiene una coleccion, _restaurantes_
## 2. ¿Cuántos documentos hay en cada colección? ¿Cuánto pesa cada colección?
- Restaurantes: 25359 documentos, 4.10KB
## 3. ¿Cuántos índices en cada colección? ¿Cuánto espacio ocupan los índices de cada colección?
- Restaurantes: 1 indice, 4.1kBs

# Ej 1
## 1.
``` js
db.restaurantes.find(
    {},
    {
        restaurante_id : 1, 
        nombre : 1, 
        barrio: 1, 
        tipo_cocina: 1,
        _id: 0
    }
)
```
## 2. 
``` js
db.restaurantes.find(
    {
      nombre: /Back/
    },
    {
        restaurante_id : 1, 
        nombre : 1, 
        barrio: 1, 
        tipo_cocina: 1,
    }
).limit(3)
```
## 3.
``` js
db.restaurantes.aggregate( [
    {
        $match: { barrio: "Bronx" }
    }, 
    {
        $group: {
        _id: '$tipo_cocina',
        count: {
            $sum: 1
        }
        }
    }, 
    {
        $project: {
            _id: 0,
            tipo_cocina: '$_id',
            count: '$count'
        }
    }
] )
```
