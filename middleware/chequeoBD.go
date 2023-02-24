package middleware

import (
	"github.com/wimuweb/red-twittor/db"
	"net/http"
)

/* ChequeoBD es el middleware que me permite conocedr el estado de  la BD */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ChequeoConexion() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
