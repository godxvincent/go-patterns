package execution

import (
	"log"

	"godxvincent.com/go-patterns/patterns/creational/builder/builders"
)

func Execute() {
	bicycleBuilder := builders.NewGTBicycleBuilder()
	mountainBicycle, err := bicycleBuilder.WithColour("red").WithHeight(29).Build()
	if err != nil {
		log.Println(err)
	}
	log.Println("This is the object created <<", mountainBicycle, ">>")
}
