# DDD Demo (Citas, Misiones, Catálogos)

Este proyecto está diseñado para **enseñanza** de DDD en Go.

## Ideas clave

- Cada contexto vive en su carpeta (`internal/citas`, `internal/misiones`, `internal/catalogos`).
- Cada contexto está separado por capas:
  - `domain`: reglas de negocio, agregados y eventos de dominio.
  - `application`: casos de uso (orquestan dominio, repositorio y publicación de eventos).
  - `infrastructure`: implementación técnica (Postgres).
  - `interfaces/http`: adaptador REST (entrada/salida JSON).

## Flujos implementados en esta iteración

- Citas: `POST /citas` -> crea una cita y emite `CitaCreada`.
- Misiones: `POST /misiones` -> crea misión desde cero y emite `MisionCreadaDesdeCero`.
- Catálogos: `POST /catalogos/usuarios` -> crea usuario y emite `UsuarioCreado`.

## Por qué hay métodos/eventos no implementados

El objetivo es mostrar **estructura y patrones** sin construir un sistema completo.
Por eso existen contratos/eventos declarados que sirven como guía para extensiones.
