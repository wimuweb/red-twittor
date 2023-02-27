package routers

import (
	"encoding/json"
	"github.com/wimuweb/red-twittor/db"
	"net/http"
)

/* VerPerfil  permite extraer los valores del perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := db.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
