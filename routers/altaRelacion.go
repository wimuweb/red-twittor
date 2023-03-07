package routers

import (
	"github.com/wimuweb/red-twittor/db"
	"github.com/wimuweb/red-twittor/models"
	"net/http"
)

/* AltaRelacion realiza el registro de la relacion entre usuarios */
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := db.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al  insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado  insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
