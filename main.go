package main

import (
	"github.com/wimuweb/red-twittor/db"
	"github.com/wimuweb/red-twittor/handlers"
	"log"
)

func main() {
	if db.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
