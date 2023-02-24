package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*MongoCN es el objeto de conexion a la base de datos */
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://wimuweb:4Rkw5aC76SyNZ2BD@miclustercafewi.739vg.mongodb.net/red-twittor")

/*ConectarDB es la funcion que me permita conectar la BD */
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client

	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/*ChequeoConexion es el Ping a la base de datos */
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {

		return 0
	}
	return 1
}
