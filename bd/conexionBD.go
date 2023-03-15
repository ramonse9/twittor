package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*MongoCN es el onjeto de conexion a la BD*/
var MongoCN = ConectarBD()

// var clientOptions = options.Client().ApplyURI("mongodb+srv://pokemondb_user:JCTu1et2U89eBKw@pokemoncluster.ywymhby.mongodb.net/?retryWrites=true&w=majority")
var clientOptions = options.Client().ApplyURI("mongodb+srv://twittor_user:o7VR9R4ByEDRDUri@pokemoncluster.ywymhby.mongodb.net/?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite copnectar la BD*/
func ConectarBD() *mongo.Client {
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

/*chequeConnection es el Ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
