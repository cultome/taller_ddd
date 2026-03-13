-- Esquema mínimo para la demo DDD.
-- Incluye únicamente tablas necesarias para los flujos implementados.

CREATE TABLE IF NOT EXISTS citas (
    id TEXT PRIMARY KEY,
    invitador_id TEXT NOT NULL,
    invitado_nombre TEXT NOT NULL,
    fecha_programada TIMESTAMPTZ NOT NULL,
    estado TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS misiones (
    id TEXT PRIMARY KEY,
    nombre TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    estado TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS catalogos_usuarios (
    id TEXT PRIMARY KEY,
    nombre TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
