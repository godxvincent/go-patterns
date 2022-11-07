package execution

import (
	"log"

	"godxvincent.com/go-patterns/spkg-patterns/builder/model"
)

func newGTBicycle() model.Bicycle {
	bicycleBuilder := model.GTBuilder{}
	mountainBicycle, err := bicycleBuilder.WithColour("red").WithHeight(29).Build()
	if err != nil {
		log.Println(err)
	}
	return mountainBicycle
}

func Execute() {
	bicycle := newGTBicycle()
	log.Println("This is the object created <<", bicycle, ">>")
}
