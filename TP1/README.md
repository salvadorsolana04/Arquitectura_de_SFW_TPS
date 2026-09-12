# TPs — De monolito a microservicios (e-commerce)

Consigna original: [`monolito`](monolito) es la base de partida (copiada sin
modificaciones del repo de la materia). El trabajo entregado está en
[`microservicios`](microservicios).

## Cómo correrlo

RabbitMQ debe estar activo (por ejemplo con el `docker compose` de `CLASE_4`
del repo de la materia) para probar la publicación del evento.

```bash
cd microservicios/clientes
go mod tidy
go run .
```

```bash
cd microservicios/pedidos
go mod tidy
go run .
```

`clientes` queda en `http://localhost:8081` y `pedidos` en
`http://localhost:8082`.

## Endpoints

| Servicio  | Método | Endpoint         |
| --------- | ------ | ---------------- |
| clientes  | POST   | `/clientes`       |
| clientes  | GET    | `/clientes/:id`   |
| pedidos   | GET    | `/productos`      |
| pedidos   | POST   | `/pedidos`        |
