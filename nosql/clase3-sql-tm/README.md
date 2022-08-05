# C3 - NoSQL TM | Práctica 1

Estas preguntas fueron respondidas utilizando la interfaz gráfica de **Compass**.

1. ¿Cuántas colecciones tiene la base de datos?

    Tiene **una** colección
    
    >```db.stats()```
    
    ```
    {
        db: 'sample_restaurants',
        collections: 1,
        views: 0,
        objects: 25359,
        avgObjSize: 439.33025750226744,
        dataSize: 11140976,
        storageSize: 4096,
        indexes: 1,
        indexSize: 4096,
        totalSize: 8192,
        scaleFactor: 1,
        fsUsedSize: 79363346432,
        fsTotalSize: 494384795648,
        ok: 1
    }
    ```
        
2. ¿Cuántos documentos hay en cada colección? ¿Cuánto pesa cada colección?
    >```db.restaurants.count()```
    
    > 25359

3. ¿Cuántos índices en cada colección? ¿Cuánto espacio ocupan los índices de cada colección?
Tiene **un** índice y ocupa 4.1 Kb
    
    >```indexes: 1,```
    
    >```indexSize: 4096,```

 
4. Traer un documento de ejemplo de cada colección. db.collection.find(...).pretty() nos da un formato más legible.
    
    >```db.restaurants.find().pretty()```

    <details>
    <summary>Resultado</summary>
    
    
        { _id: ObjectId("5eb3d668b31de5d588f4294f"),
          direccion: 
           { edificio: '625',
             coord: [ -73.990494, 40.7569545 ],
             calle: '8 Avenue',
             codigo_postal: '10018' },
          barrio: 'Manhattan',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-06-09T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2014-01-10T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2012-12-07T00:00:00.000Z, grado: 'A', puntaje: 4 },
             { date: 2011-12-13T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2011-09-09T00:00:00.000Z, grado: 'A', puntaje: 13 } ],
          nombre: 'Cafe Metro',
          restaurante_id: '40363298' }
        { _id: ObjectId("5eb3d668b31de5d588f42930"),
          direccion: 
           { edificio: '8825',
             coord: [ -73.8803827, 40.7643124 ],
             calle: 'Astoria Boulevard',
             codigo_postal: '11369' },
          barrio: 'Queens',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-11-15T00:00:00.000Z, grado: 'Z', puntaje: 38 },
             { date: 2014-05-02T00:00:00.000Z, grado: 'A', puntaje: 10 },
             { date: 2013-03-02T00:00:00.000Z, grado: 'A', puntaje: 7 },
             { date: 2012-02-10T00:00:00.000Z, grado: 'A', puntaje: 13 } ],
          nombre: 'Brunos On The Boulevard',
          restaurante_id: '40356151' }
        { _id: ObjectId("5eb3d668b31de5d588f42955"),
          direccion: 
           { edificio: '464',
             coord: [ -73.9791458, 40.744328 ],
             calle: '3 Avenue',
             codigo_postal: '10016' },
          barrio: 'Manhattan',
          tipo_cocina: 'Pizza',
          grados: 
           [ { date: 2014-08-05T00:00:00.000Z, grado: 'A', puntaje: 3 },
             { date: 2014-03-06T00:00:00.000Z, grado: 'A', puntaje: 11 },
             { date: 2013-07-09T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2013-01-30T00:00:00.000Z, grado: 'A', puntaje: 4 },
             { date: 2012-01-05T00:00:00.000Z, grado: 'A', puntaje: 2 },
             { date: 2011-09-26T00:00:00.000Z, grado: 'A', puntaje: 0 } ],
          nombre: 'Domino\'S Pizza',
          restaurante_id: '40363644' }
        { _id: ObjectId("5eb3d668b31de5d588f4295b"),
          direccion: 
           { edificio: '1031',
             coord: [ -73.9075537, 40.6438684 ],
             calle: 'East   92 Street',
             codigo_postal: '11236' },
          barrio: 'Brooklyn',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-02-05T00:00:00.000Z, grado: 'A', puntaje: 0 },
             { date: 2013-01-29T00:00:00.000Z, grado: 'A', puntaje: 3 },
             { date: 2011-12-08T00:00:00.000Z, grado: 'A', puntaje: 10 } ],
          nombre: 'Sonny\'S Heros',
          restaurante_id: '40363744' }
        { _id: ObjectId("5eb3d668b31de5d588f42978"),
          direccion: 
           { edificio: '2602',
             coord: [ -73.95443709999999, 40.5877993 ],
             calle: 'East   15 Street',
             codigo_postal: '11235' },
          barrio: 'Brooklyn',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-05-14T00:00:00.000Z, grado: 'A', puntaje: 11 },
             { date: 2013-04-27T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2012-11-23T00:00:00.000Z, grado: 'B', puntaje: 27 },
             { date: 2012-03-14T00:00:00.000Z, grado: 'B', puntaje: 17 },
             { date: 2011-07-14T00:00:00.000Z, grado: 'B', puntaje: 21 } ],
          nombre: 'Towne Cafe',
          restaurante_id: '40364681' }
        { _id: ObjectId("5eb3d668b31de5d588f4299a"),
          direccion: 
           { edificio: '72',
             coord: [ -73.92506, 40.8275556 ],
             calle: 'East  161 Street',
             codigo_postal: '10451' },
          barrio: 'Bronx',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-04-15T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2013-11-14T00:00:00.000Z, grado: 'A', puntaje: 4 },
             { date: 2013-07-29T00:00:00.000Z, grado: 'A', puntaje: 10 },
             { date: 2012-12-31T00:00:00.000Z, grado: 'B', puntaje: 15 },
             { date: 2012-05-30T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2012-01-09T00:00:00.000Z, grado: 'A', puntaje: 10 },
             { date: 2011-08-15T00:00:00.000Z, grado: 'C', puntaje: 37 } ],
          nombre: 'Yankee Tavern',
          restaurante_id: '40365499' }
        { _id: ObjectId("5eb3d668b31de5d588f429b0"),
          direccion: 
           { edificio: '416',
             coord: [ -73.98586209999999, 40.67017250000001 ],
             calle: '5 Avenue',
             codigo_postal: '11215' },
          barrio: 'Brooklyn',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-12-04T00:00:00.000Z, grado: 'B', puntaje: 22 },
             { date: 2014-04-19T00:00:00.000Z, grado: 'A', puntaje: 3 },
             { date: 2013-02-14T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2012-01-12T00:00:00.000Z, grado: 'A', puntaje: 9 } ],
          nombre: 'Fifth Avenue Bingo',
          restaurante_id: '40366109' }
        { _id: ObjectId("5eb3d668b31de5d588f429d8"),
          direccion: 
           { edificio: '4',
             coord: [ -74.00065800000002, 40.735114 ],
             calle: 'Charles Street',
             codigo_postal: '10014' },
          barrio: 'Manhattan',
          tipo_cocina: 'Spanish',
          grados: 
           [ { date: 2014-04-16T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2013-10-28T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2012-07-12T00:00:00.000Z, grado: 'A', puntaje: 2 },
             { date: 2012-02-28T00:00:00.000Z, grado: 'A', puntaje: 5 },
             { date: 2011-10-06T00:00:00.000Z, grado: 'A', puntaje: 2 },
             { date: 2011-06-02T00:00:00.000Z, grado: 'A', puntaje: 7 } ],
          nombre: 'El Charro Espanol',
          restaurante_id: '40366987' }
        { _id: ObjectId("5eb3d668b31de5d588f429f1"),
          direccion: 
           { edificio: '4933',
             coord: [ -73.9215284, 40.8678204 ],
             calle: 'Broadway',
             codigo_postal: '10034' },
          barrio: 'Manhattan',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-07-24T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2013-07-15T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2013-02-20T00:00:00.000Z, grado: 'A', puntaje: 6 },
             { date: 2012-08-31T00:00:00.000Z, grado: 'A', puntaje: 11 },
             { date: 2012-04-02T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2011-11-16T00:00:00.000Z, grado: 'A', puntaje: 12 } ],
          nombre: 'Capitol Restaurant',
          restaurante_id: '40367677' }
        { _id: ObjectId("5eb3d668b31de5d588f429f2"),
          direccion: 
           { edificio: '48',
             coord: [ -73.977035, 40.762307 ],
             calle: 'West   55 Street',
             codigo_postal: '10019' },
          barrio: 'Manhattan',
          tipo_cocina: 'French',
          grados: 
           [ { date: 2014-05-15T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2013-12-16T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2013-05-21T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2012-05-07T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2011-12-12T00:00:00.000Z, grado: 'A', puntaje: 11 },
             { date: 2011-07-25T00:00:00.000Z, grado: 'A', puntaje: 5 } ],
          nombre: 'La Bonne Soupe Bistro',
          restaurante_id: '40367715' }
        { _id: ObjectId("5eb3d668b31de5d588f429f7"),
          direccion: 
           { edificio: '2080',
             coord: [ -73.98185529999999, 40.7782266 ],
             calle: 'Broadway',
             codigo_postal: '10023' },
          barrio: 'Manhattan',
          tipo_cocina: 'Hotdogs',
          grados: 
           [ { date: 2014-02-18T00:00:00.000Z, grado: 'A', puntaje: 7 },
             { date: 2013-08-23T00:00:00.000Z, grado: 'A', puntaje: 8 },
             { date: 2013-01-16T00:00:00.000Z, grado: 'B', puntaje: 20 },
             { date: 2012-01-05T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2011-09-13T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2011-04-28T00:00:00.000Z, grado: 'A', puntaje: 7 } ],
          nombre: 'Gray\'S Papaya',
          restaurante_id: '40367766' }
        { _id: ObjectId("5eb3d668b31de5d588f42a19"),
          direccion: 
           { edificio: '59',
             coord: [ -74.00340299999999, 40.733235 ],
             calle: 'Grove Street',
             codigo_postal: '10014' },
          barrio: 'Manhattan',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-10-15T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2014-04-22T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2013-07-30T00:00:00.000Z, grado: 'B', puntaje: 23 },
             { date: 2012-06-14T00:00:00.000Z, grado: 'C', puntaje: 29 },
             { date: 2011-12-15T00:00:00.000Z, grado: 'A', puntaje: 12 } ],
          nombre: 'Maries Crisis Cafe',
          restaurante_id: '40368581' }
        { _id: ObjectId("5eb3d668b31de5d588f42a21"),
          direccion: 
           { edificio: '840',
             coord: [ -73.970668, 40.751453 ],
             calle: '2 Avenue',
             codigo_postal: '10017' },
          barrio: 'Manhattan',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-03-31T00:00:00.000Z, grado: 'A', puntaje: 11 },
             { date: 2013-03-27T00:00:00.000Z, grado: 'A', puntaje: 5 },
             { date: 2012-10-15T00:00:00.000Z, grado: 'A', puntaje: 3 },
             { date: 2012-05-15T00:00:00.000Z, grado: 'A', puntaje: 12 } ],
          nombre: 'Palm Too',
          restaurante_id: '40369017' }
        { _id: ObjectId("5eb3d668b31de5d588f42a2f"),
          direccion: 
           { edificio: '395',
             coord: [ -73.9808063, 40.6895078 ],
             calle: 'Flatbush Avenue Extension',
             codigo_postal: '11201' },
          barrio: 'Brooklyn',
          tipo_cocina: 'Hamburgers',
          grados: 
           [ { date: 2014-04-07T00:00:00.000Z, grado: 'A', puntaje: 10 },
             { date: 2013-04-04T00:00:00.000Z, grado: 'A', puntaje: 7 },
             { date: 2012-04-06T00:00:00.000Z, grado: 'A', puntaje: 5 },
             { date: 2011-04-13T00:00:00.000Z, grado: 'A', puntaje: 10 } ],
          nombre: 'Mcdonald\'S',
          restaurante_id: '40369535' }
        { _id: ObjectId("5eb3d668b31de5d588f42a41"),
          direccion: 
           { edificio: '1011',
             coord: [ -73.9623333, 40.7757194 ],
             calle: 'Madison Avenue',
             codigo_postal: '10075' },
          barrio: 'Manhattan',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-05-02T00:00:00.000Z, grado: 'A', puntaje: 10 },
             { date: 2013-10-21T00:00:00.000Z, grado: 'B', puntaje: 15 },
             { date: 2013-04-19T00:00:00.000Z, grado: 'B', puntaje: 0 },
             { date: 2012-11-21T00:00:00.000Z, grado: 'C', puntaje: 40 },
             { date: 2012-04-09T00:00:00.000Z, grado: 'B', puntaje: 17 } ],
          nombre: 'Viand Cafe',
          restaurante_id: '40369753' }
        { _id: ObjectId("5eb3d668b31de5d588f42a5e"),
          direccion: 
           { edificio: '138140',
             coord: [ -73.9574128, 40.7701235 ],
             calle: 'East   74 Street',
             codigo_postal: '10021' },
          barrio: 'Manhattan',
          tipo_cocina: 'Italian',
          grados: 
           [ { date: 2014-05-02T00:00:00.000Z, grado: 'A', puntaje: 9 },
             { date: 2013-04-12T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2012-04-03T00:00:00.000Z, grado: 'A', puntaje: 11 } ],
          nombre: 'Cucina Vivolo',
          restaurante_id: '40370497' }
        { _id: ObjectId("5eb3d668b31de5d588f42a62"),
          direccion: 
           { edificio: '16',
             coord: [ -73.986685, 40.73756400000001 ],
             calle: 'Gramercy Park South',
             codigo_postal: '10003' },
          barrio: 'Manhattan',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-11-13T00:00:00.000Z, grado: 'Z', puntaje: 25 },
             { date: 2014-04-29T00:00:00.000Z, grado: 'A', puntaje: 7 },
             { date: 2013-09-17T00:00:00.000Z, grado: 'C', puntaje: 35 },
             { date: 2013-02-19T00:00:00.000Z, grado: 'A', puntaje: 10 },
             { date: 2012-06-12T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2011-12-22T00:00:00.000Z, grado: 'A', puntaje: 13 } ],
          nombre: 'The Players Club',
          restaurante_id: '40370507' }
        { _id: ObjectId("5eb3d668b31de5d588f42a67"),
          direccion: 
           { edificio: '1745',
             coord: [ -74.146976, 40.625178 ],
             calle: 'Forest Avenue',
             codigo_postal: '10303' },
          barrio: 'Staten Island',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-11-15T00:00:00.000Z, grado: 'A', puntaje: 8 },
             { date: 2014-06-24T00:00:00.000Z, grado: 'A', puntaje: 10 },
             { date: 2014-01-08T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2012-12-07T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2012-06-23T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2012-01-24T00:00:00.000Z, grado: 'A', puntaje: 9 } ],
          nombre: 'Perkins Family Restaurant & Bakery',
          restaurante_id: '40370910' }
        { _id: ObjectId("5eb3d668b31de5d588f42a68"),
          direccion: 
           { edificio: '1701',
             coord: [ -73.96123879999999, 40.635193 ],
             calle: 'Foster Avenue',
             codigo_postal: '11230' },
          barrio: 'Brooklyn',
          tipo_cocina: 'Italian',
          grados: 
           [ { date: 2014-05-13T00:00:00.000Z, grado: 'A', puntaje: 12 },
             { date: 2013-05-09T00:00:00.000Z, grado: 'A', puntaje: 7 },
             { date: 2012-10-17T00:00:00.000Z, grado: 'B', puntaje: 27 },
             { date: 2012-08-28T00:00:00.000Z, grado: 'P', puntaje: 4 },
             { date: 2012-03-07T00:00:00.000Z, grado: 'A', puntaje: 12 } ],
          nombre: 'Mama Lucia',
          restaurante_id: '40370994' }
        { _id: ObjectId("5eb3d668b31de5d588f42a75"),
          direccion: 
           { edificio: '2055',
             coord: [ -74.1321, 40.61266000000001 ],
             calle: 'Victory Boulevard',
             codigo_postal: '10314' },
          barrio: 'Staten Island',
          tipo_cocina: 'American',
          grados: 
           [ { date: 2014-11-06T00:00:00.000Z, grado: 'B', puntaje: 25 },
             { date: 2014-05-06T00:00:00.000Z, grado: 'B', puntaje: 20 },
             { date: 2013-01-26T00:00:00.000Z, grado: 'A', puntaje: 13 },
             { date: 2011-12-17T00:00:00.000Z, grado: 'A', puntaje: 7 } ],
          nombre: 'Schaffer\'S Tavern',
          restaurante_id: '40371771' }
        Type "it" for more
        ```

    </details>
    
        
        
        
5. Para cada colección, listar los campos a nivel raíz (ignorar campos dentro de documentos anidados) y sus tipos de datos.
    
    >```Object.keys(db.restaurants.findOne())```
    
    ```
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
    

# **Ejercicio 1: SQL**

Usando Mongo Shell. Colección restaurantes se requiere:

1. Devolver restaurante_id, nombre, barrio y tipo_cocina pero excluyendo _id para un documento (el primero).
    
    > ```db.restaurants.find({},{_id: 0, restaurante_id: 1, status: 1, nombre: 1, status: 1, barrio: 1, status :1 ,tipo_cocina: 1, status: 1 })```

    <details>
    <summary>Resultado</summary>
            
            { barrio: 'Manhattan',
            tipo_cocina: 'American',
            nombre: 'Cafe Metro',
            restaurante_id: '40363298' }
            { barrio: 'Queens',
            tipo_cocina: 'American',
            nombre: 'Brunos On The Boulevard',
            restaurante_id: '40356151' }
            { barrio: 'Manhattan',
            tipo_cocina: 'Pizza',
            nombre: 'Domino\'S Pizza',
            restaurante_id: '40363644' }
            { barrio: 'Brooklyn',
            tipo_cocina: 'American',
            nombre: 'Sonny\'S Heros',
            restaurante_id: '40363744' }
            { barrio: 'Brooklyn',
            tipo_cocina: 'American',
            nombre: 'Towne Cafe',
            restaurante_id: '40364681' }
            { barrio: 'Bronx',
            tipo_cocina: 'American',
            nombre: 'Yankee Tavern',
            restaurante_id: '40365499' }
            { barrio: 'Brooklyn',
            tipo_cocina: 'American',
            nombre: 'Fifth Avenue Bingo',
            restaurante_id: '40366109' }
            { barrio: 'Manhattan',
            tipo_cocina: 'Spanish',
            nombre: 'El Charro Espanol',
            restaurante_id: '40366987' }
            { barrio: 'Manhattan',
            tipo_cocina: 'American',
            nombre: 'Capitol Restaurant',
            restaurante_id: '40367677' }
            { barrio: 'Manhattan',
            tipo_cocina: 'French',
            nombre: 'La Bonne Soupe Bistro',
            restaurante_id: '40367715' }
            { barrio: 'Manhattan',
            tipo_cocina: 'Hotdogs',
            nombre: 'Gray\'S Papaya',
            restaurante_id: '40367766' }
            { barrio: 'Manhattan',
            tipo_cocina: 'American',
            nombre: 'Maries Crisis Cafe',
            restaurante_id: '40368581' }
            { barrio: 'Manhattan',
            tipo_cocina: 'American',
            nombre: 'Palm Too',
            restaurante_id: '40369017' }
            { barrio: 'Brooklyn',
            tipo_cocina: 'Hamburgers',
            nombre: 'Mcdonald\'S',
            restaurante_id: '40369535' }
            { barrio: 'Manhattan',
            tipo_cocina: 'American',
            nombre: 'Viand Cafe',
            restaurante_id: '40369753' }
            { barrio: 'Manhattan',
            tipo_cocina: 'Italian',
            nombre: 'Cucina Vivolo',
            restaurante_id: '40370497' }
            { barrio: 'Manhattan',
            tipo_cocina: 'American',
            nombre: 'The Players Club',
            restaurante_id: '40370507' }
            { barrio: 'Staten Island',
            tipo_cocina: 'American',
            nombre: 'Perkins Family Restaurant & Bakery',
            restaurante_id: '40370910' }
            { barrio: 'Brooklyn',
            tipo_cocina: 'Italian',
            nombre: 'Mama Lucia',
            restaurante_id: '40370994' }
            { barrio: 'Staten Island',
            tipo_cocina: 'American',
            nombre: 'Schaffer\'S Tavern',
            restaurante_id: '40
        ``` json
    </details>

        
2. Devolver restaurante_id, nombre, barrio y tipo_cocina para los primeros 3 restaurantes que contengan 'Bake' en alguna parte de su nombre.
`> db.restaurants.find({},{_id: 0, restaurante_id: 1, status: 1, nombre: 1, status: 1, barrio: 1, status :1 ,tipo_cocina: 1, status: 1 }).limit(3)`
    
    ```
    { barrio: 'Manhattan',
      tipo_cocina: 'American',
      nombre: 'Cafe Metro',
      restaurante_id: '40363298' }
    { barrio: 'Queens',
      tipo_cocina: 'American',
      nombre: 'Brunos On The Boulevard',
      restaurante_id: '40356151' }
    { barrio: 'Manhattan',
      tipo_cocina: 'Pizza',
      nombre: 'Domino\'S Pizza',
      restaurante_id: '40363644' }
    ```
    
3. Contar los restaurantes de comida (tipo_cocina) china (*Chinese*) o tailandesa (*Thai*) del barrio (barrio) Bronx. Consultar [or versus in](https://docs.mongodb.com/manual/reference/operator/query/or/#-or-versus--in).
    
    > ```db.restaurants.find( { $or:[ {tipo_cocina: "Chinese"}, {tipo_cocina: "Thai" }] }).count()```
      
    > 2703

# **Ejercicio 2: NoSQL**

1. Traer 3 restaurantes que hayan recibido al menos una calificación de grado 'A' con puntaje mayor a 20. Una misma calificación debe cumplir con ambas condiciones simultáneamente; investigar el operador [elemMatch](https://docs.mongodb.com/manual/reference/operator/query/elemMatch/).
    
    ```bash
    db.restaurants.count(
       { tipo_cocina: { $in: [ "Chinese", "Thai" ] }, barrio: "Bronx" }
    )
    ```
    
    > 325
    > 
2. ¿A cuántos documentos les faltan las coordenadas geográficas? En otras palabras, revisar si el tamaño de direccion.coord es 0 y contar.

```bash
db.restaurants.count(
   { "direccion.coord": { $size: 0 } }
)
```

> 2
> 
3. Devolver nombre, barrio, tipo_cocina y grados para los primeros 3 restaurantes; de cada documento **solo la última calificación**. Ver el operador [slice](https://docs.mongodb.com/manual/tutorial/project-fields-from-query-results/#project-specific-array-elements-in-the-returned-array).

```bash
db.restaurants.find(
   { },
   { nombre: 1, barrio: 1, tipo_cocina: 1, grados: { $slice: -1 } }
).limit(3)
```

```
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