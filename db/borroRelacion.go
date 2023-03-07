package db

import (
	"context"
	"github.com/wimuweb/red-twittor/models"
	"time"
)

/*BorroRelacion borra la relacion en la BD */
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("red-twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil

}
