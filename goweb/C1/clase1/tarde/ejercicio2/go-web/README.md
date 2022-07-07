## Como realizar solicitudes con gin en el projecto

1. Realizar filtrado
```curl
curl -X GET http://localhost:8000/filtrartransaction?id=1&codigo=abc123&moneda=peso&monto=100&emisor=Martín&receptor=Luisa
```

2. Get transición /:id
```curl
curl -X GET http://localhost:8000/transactions/1 
curl -X GET http://localhost:8000/transactions/1 -I
```
