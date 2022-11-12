package builders

import (
	validator "github.com/go-playground/validator/v10"
	"godxvincent.com/go-patterns/patterns/creational/builder/model"
)

type GTBuilder struct {
	height int
	colour string
}

func (gtbuilder *GTBuilder) WithHeight(height int) model.IBicycleBuilder {
	gtbuilder.height = height
	return gtbuilder
}

func (gtbuilder *GTBuilder) WithColour(colour string) model.IBicycleBuilder {
	gtbuilder.colour = colour
	return gtbuilder
}

func (gtbuilder *GTBuilder) Build() (model.Bicycle, error) {
	gtbicycle := model.Bicycle{
		Make:   "GT",
		Model:  "Avalanche",
		Height: gtbuilder.height,
		Colour: gtbuilder.colour,
	}
	validator := validator.New()
	err := validator.Struct(gtbicycle)
	if err != nil {
		gtbicycle = model.Bicycle{}
	}
	return gtbicycle, err
}

func NewGTBicycleBuilder() GTBuilder {
	bicycleBuilder := GTBuilder{}
	return bicycleBuilder
}