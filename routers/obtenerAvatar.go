package routers

import (
	"github.com/wimuweb/red-twittor/db"
	"io"
	"net/http"
	"os"
)

/*ObtenerAvatar envia el Avatar al HTTP */
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id ! ", http.StatusBadRequest)
		return
	}
	perfil, err := db.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado ! ", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada ! ", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! ", http.StatusBadRequest)
	}

}
