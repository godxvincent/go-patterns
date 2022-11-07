package execution

import (
	"log"

	"godxvincent.com/go-patterns/spkg-patterns/singleton/model"
)

func Execute() {
	log.Println("Se crea nueva instancia deconexión")
	connection := model.CreateConnection()
	log.Println("Se consulta el estado de la conexión creada <<", connection.GetState(), ">>")
	log.Println("Se setea el estado de la conexión como iniciada")
	connection.SetState("Iniciada")
	log.Println("Estado de la conexión tras setearla <<", connection.GetState(), ">>")
	log.Println("Se intenta crea una nueva instancia de la conexión")
	connection = model.CreateConnection()
	log.Println("Estado de la conexión recreada <<", connection.GetState(), ">>")
}
