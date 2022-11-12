package model

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
