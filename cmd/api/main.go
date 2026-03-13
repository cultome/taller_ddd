package main

import (
	"log"
	"net/http"

	catalogosapp "aliado_ddd/internal/catalogos/application"
	catalogospg "aliado_ddd/internal/catalogos/infrastructure/postgres"
	catalogoshttp "aliado_ddd/internal/catalogos/interfaces/http"
	citasapp "aliado_ddd/internal/citas/application"
	citaspg "aliado_ddd/internal/citas/infrastructure/postgres"
	citashttp "aliado_ddd/internal/citas/interfaces/http"
	misionesapp "aliado_ddd/internal/misiones/application"
	misionespg "aliado_ddd/internal/misiones/infrastructure/postgres"
	misioneshttp "aliado_ddd/internal/misiones/interfaces/http"
	"aliado_ddd/internal/shared/infrastructure/events"
	sharedpg "aliado_ddd/internal/shared/infrastructure/postgres"
)

func main() {
	db, err := sharedpg.OpenDB()
	if err != nil {
		log.Fatalf("db error: %v", err)
	}
	defer db.Close()

	dispatcher := events.NewInMemoryDispatcher()

	// Wiring de dependencias por contexto (DDD):
	// infra -> application -> interfaces
	citaRepo := citaspg.NewCitaRepository(db)
	createCita := citasapp.NewCreateCitaService(citaRepo, dispatcher)
	citaHandler := citashttp.NewHandler(createCita)

	misionRepo := misionespg.NewMisionRepository(db)
	createMision := misionesapp.NewCreateMisionService(misionRepo, dispatcher)
	misionHandler := misioneshttp.NewHandler(createMision)

	usuarioRepo := catalogospg.NewUsuarioRepository(db)
	createUsuario := catalogosapp.NewCreateUsuarioService(usuarioRepo, dispatcher)
	usuarioHandler := catalogoshttp.NewHandler(createUsuario)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// Un flujo implementado por contexto.
	mux.HandleFunc("POST /citas", citaHandler.CreateCita)
	mux.HandleFunc("POST /misiones", misionHandler.CreateMision)
	mux.HandleFunc("POST /catalogos/usuarios", usuarioHandler.CreateUsuario)

	log.Println("API listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
