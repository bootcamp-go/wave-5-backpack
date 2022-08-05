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
