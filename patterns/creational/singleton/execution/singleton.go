package execution

import (
	"log"

	"godxvincent.com/go-patterns/patterns/creational/singleton/model"
)

func Execute() {
	log.Println("New connection instance is created")
	connection := model.CreateConnection()
	log.Println("Getting connection status <<", connection.GetState(), ">>")
	log.Println("Setting new connection status")
	connection.SetState("Started")
	log.Println("Getting connection status after being setted <<", connection.GetState(), ">>")
	log.Println("Trying to create a new connection instance")
	connection = model.CreateConnection()
	log.Println("Getting connection status <<", connection.GetState(), ">>")
}
