package main

import (
	"github.com/ramonse9/twittor/bd"
	"github.com/ramonse9/twittor/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
