# Arquitectura — TP1

## Antes

```
cliente -> monolito :8080 (clientes + productos + pedidos, todo en un proceso)
```

## Después

```
cliente -> clientes :8081  (alta y consulta de clientes)
        -> pedidos  :8082  (catálogo de productos + confirmación de pedidos)

pedidos -- evento pedido.confirmado --> RabbitMQ (cola "pedidos-confirmados") --> logística (conceptual, no implementado)
```

## Decisiones de diseño

- **Separación por dominio**: `clientes` y `pedidos` son módulos Go
  independientes (cada uno con su propio `go.mod`), sin dependencias en tiempo
  de compilación entre sí. Se pueden levantar, compilar y desplegar por
  separado.
- **Capas**: cada microservicio sigue `controller -> service -> repository`.
  Los controllers solo traducen HTTP <-> dominio, los services concentran la
  lógica de negocio, y los repositories son la única capa que conoce cómo se
  guardan los datos (en este caso, en memoria).
- **Catálogo de productos en `pedidos`**: la consigna pide que `pedidos`
  implemente `GET /productos`, así que el catálogo vive ahí y no en
  `clientes`.
- **Caché de productos**: `ProductoService` cachea en memoria el resultado de
  `ProductoRepository.ListarTodos()` por 30 segundos (`cacheTTL`), evitando
  golpear el repositorio en cada request. `ProductoRepository` simula una
  consulta costosa con un `time.Sleep`, para que la caché tenga un efecto
  real y medible.
- **Publicación del evento**: `PedidoService` depende de la interfaz
  `Publicador` (no de la implementación concreta de RabbitMQ), así que la
  lógica de negocio queda desacoplada de la mensajería. `RabbitMQPublisher`
  intenta conectarse al arrancar; si RabbitMQ no está disponible, el
  microservicio sigue funcionando igual y solo loguea que no pudo publicar
  (para no bloquear el desarrollo local si todavía no se levantó RabbitMQ).
- **IDs en memoria**: tanto clientes como pedidos generan IDs incrementales
  (`C-1`, `PED-1`, ...) protegidos por mutex, ya que no hay base de datos real
  (fuera del alcance del TP).
