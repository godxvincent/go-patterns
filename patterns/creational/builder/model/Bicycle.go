package model

import validator "github.com/go-playground/validator/v10"

// The interface define what methods the concret builder should implement
type IBicycleBuilder interface {
	WithHeight(int) IBicycleBuilder
	WithColour(string) IBicycleBuilder
	Build() (Bicycle, error)
}

// The class or struct in case of golang that define the attributes.
type Bicycle struct {
	Make   string `validate:"required"`
	Model  string `validate:"required"`
	Height int    `validate:"required"`
	Colour string `validate:"required"`
}

type GTBuilder struct {
	height int
	colour string
}

func (gtbuilder *GTBuilder) WithHeight(height int) IBicycleBuilder {
	gtbuilder.height = height
	return gtbuilder
}

func (gtbuilder *GTBuilder) WithColour(colour string) IBicycleBuilder {
	gtbuilder.colour = colour
	return gtbuilder
}

func (gtbuilder *GTBuilder) Build() (Bicycle, error) {
	gtbicycle := Bicycle{
		Make:   "GT",
		Model:  "Avalanche",
		Height: gtbuilder.height,
		Colour: gtbuilder.colour,
	}
	validator := validator.New()
	err := validator.Struct(gtbicycle)
	if err != nil {
		gtbicycle = Bicycle{}
	}
	return gtbicycle, err
}
