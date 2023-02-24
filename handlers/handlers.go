package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wimuweb/red-twittor/middleware"
	"github.com/wimuweb/red-twittor/routers"
	"log"
	"net/http"
	"os"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor */
func Manejadores() {
	router := mux.NewRouter()

	/* ---- Rutas de mi backend --- */
	// Registro de usuario
	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8805"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
