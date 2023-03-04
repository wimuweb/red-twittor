package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wimuweb/red-twittor/middleware"
	"github.com/wimuweb/red-twittor/routers"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor */
func Manejadores() {
	router := mux.NewRouter()

	/* ---- Rutas de mi backend --- */
	// Registro de usuario
	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")

	//  Login de usuario
	router.HandleFunc("/login", middleware.ChequeoBD(routers.Login)).Methods("POST")

	// Ver perfil
	router.HandleFunc("/verperfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")

	// Modificar perfil
	router.HandleFunc("/modificarPerfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")

	// Grabar Tweet
	router.HandleFunc("/tweet", middleware.ChequeoBD(middleware.ValidoJWT(routers.GraboTweet))).Methods("POST")

	// Leer Tweets
	router.HandleFunc("/leoTweets", middleware.ChequeoBD(middleware.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8805"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
