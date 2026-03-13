# aliado_ddd

Demo DDD en Go con 3 contextos: Citas, Misiones y Catálogos.

## Requisitos

- Go 1.22+
- Postgres

## Variables

- `DATABASE_URL` (opcional).  
  Valor por defecto: `postgres://postgres:postgres@localhost:5432/aliado_ddd?sslmode=disable`

## Ejecutar

1. Crear base de datos `aliado_ddd`.
2. Ejecutar `migrations/001_init.sql`.
3. Ejecutar:

```bash
go mod tidy
go run ./cmd/api
```

## Endpoints

- `GET /health`
- `POST /citas`
- `POST /misiones`
- `POST /catalogos/usuarios`
