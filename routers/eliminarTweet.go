package routers

import (
	"github.com/wimuweb/red-twittor/db"
	"net/http"
)

/*EliminarTweet permite borrar un tweet determinado */
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := db.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrior un error al intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
