# Aliado DDD

Demo DDD en Go con 3 contextos: Citas, Misiones y Catálogos.

## Endpoints

- `GET /health`
- `POST /citas`
- `POST /misiones`
- `POST /catalogos/usuarios`

## Project Structure

```
├── cmd
│   └── api
│       └── main.go
├── docs
│   └── ddd-overview.md
├── go.mod
├── go.sum
├── internal
│   ├── catalogos
│   │   ├── application
│   │   │   └── create_usuario.go
│   │   ├── domain
│   │   │   └── usuario.go
│   │   ├── infrastructure
│   │   │   └── postgres
│   │   │       └── usuario_repository.go
│   │   └── interfaces
│   │       └── http
│   │           └── handler.go
│   ├── citas
│   │   ├── application
│   │   │   └── create_cita.go
│   │   ├── domain
│   │   │   └── cita.go
│   │   ├── infrastructure
│   │   │   └── postgres
│   │   │       └── cita_repository.go
│   │   └── interfaces
│   │       └── http
│   │           └── handler.go
│   ├── misiones
│   │   ├── application
│   │   │   └── create_mision.go
│   │   ├── domain
│   │   │   └── mision.go
│   │   ├── infrastructure
│   │   │   └── postgres
│   │   │       └── mision_repository.go
│   │   └── interfaces
│   │       └── http
│   │           └── handler.go
│   ├── platform
│   │   └── http
│   └── shared
│       ├── domain
│       │   ├── event.go
│       │   └── id.go
│       └── infrastructure
│           ├── events
│           │   └── dispatcher.go
│           └── postgres
│               └── db.go
├── migrations
│   └── 001_init.sql
└── README.md
```

# Componentes de DDD y su Función

1. Entidad (Entity)

Uso: Objetos que tienen una identidad única que persiste en el tiempo (un ID).
Ejemplo: Un Usuario o un Pedido. Aunque cambien su nombre o dirección, siguen siendo el mismo registro.

2. Objeto de Valor (Value Object)

Uso: Objetos sin identidad propia, definidos solo por sus atributos. Son inmutables.
Ejemplo: Dinero(cantidad, moneda) o Direccion(calle, ciudad). Si cambias la calle, es una dirección nueva, no la "misma" dirección modificada.

3. Agregado (Aggregate)

Uso: Un clúster de Entidades y Objetos de Valor que se tratan como una sola unidad de consistencia.
Ejemplo: Un Pedido (Raíz) y sus LineasDePedido. No puedes modificar una línea sin pasar por las reglas del pedido completo.

4. Repositorio (Repository - El Puerto)

Uso: Una interfaz que simula una colección de objetos en memoria. Sirve para recuperar y guardar Agregados completos.
Ejemplo: IPedidoRepository.ObtenerPorId(id).

5. Servicio de Dominio (Domain Service)

Uso: Orquesta lógica de negocio que no pertenece naturalmente a una sola Entidad o que involucra a varios Agregados. No tienen estado.
Ejemplo: TransferenciaService que coordina entre dos CuentasBancarias.

6. Servicio de Aplicación (Application Service)

Uso: El "director de orquesta". No contiene lógica de negocio, sino que coordina la infraestructura: usa repositorios, llama al dominio y guarda los cambios.
Ejemplo: CrearPedidoHandler.

7. Evento de Dominio (Domain Event)

Uso: Algo que sucedió en el dominio y que a otras partes del sistema les interesa saber.
Ejemplo: PedidoPagado, EmailCambiado.

# Resumen de Responsabilidades

| Componente             | ¿Tiene lógica de negocio? | ¿Accede a la Base de Datos? | Responsabilidad Principal                           |
| ---------------------- | ------------------------- | --------------------------- | --------------------------------------------------- |
| Entidad / VO           | SÍ                        | NO                          | Validar reglas internas y mantener su estado.       |
| Agregado               | SÍ                        | NO                          | Garantizar las invariantes de todo el grupo.        |
| Servicio de Dominio    | SÍ                        | NO (usa puertos)            | Lógica compleja que afecta a varios objetos.        |
| Repositorio            | NO                        | SÍ (vía adaptador)          | Persistencia y recuperación de Agregados.           |
| Servicio de Aplicación | NO                        | SÍ (usa el repo)            | Orquestar el flujo (Cargar -> Ejecutar -> Guardar). |
