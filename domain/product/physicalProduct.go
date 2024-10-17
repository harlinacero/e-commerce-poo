package product

import "fmt"

type PhysicalProduct struct {
	ProductBase        // Hereda de producto
	weight  float64 // Peso
	height  float64 // Altura
	width   float64 // Ancho
	length  float64 // Largo
}

// Constructor
func NewPhysicalProduct(code string, name string, description string, grossvalue float64, discPercentage float64, taxPercentage float64, stock int, weight float64, height float64, width float64, length float64) *PhysicalProduct {
	return &PhysicalProduct{
		ProductBase: ProductBase{
			code:           code,
			name:           name,
			description:    description,
			grossValue:     grossvalue,
			discPercentage: discPercentage,
			taxPercentage:  taxPercentage,
			stock:          stock,
		},
		weight: weight,
		height: height,
		width:  width,
		length: length,
	}
}

// Setters
func (pp *PhysicalProduct) SetWeight(weight float64) {
	pp.weight = weight
}

func (pp *PhysicalProduct) SetHeight(height float64) {
	pp.height = height
}

func (pp *PhysicalProduct) SetWidth(width float64) {
	pp.width = width
}

func (pp *PhysicalProduct) SetLength(length float64) {
	pp.length = length
}

// Getters
func (pp *PhysicalProduct) GetWeight() float64 {
	return pp.weight
}

func (pp *PhysicalProduct) GetHeight() float64 {
	return pp.height
}

func (pp *PhysicalProduct) GetWidth() float64 {
	return pp.width
}

func (pp *PhysicalProduct) GetLength() float64 {
	return pp.length
}

// Calcula el volumen del producto
func (pp *PhysicalProduct) GetVolume() float64 {
	return pp.height * pp.width * pp.length
}

func (pp *PhysicalProduct) ShowInfo() string {
	return fmt.Sprintln("Peso: ", pp.weight, 
	"\nAltura: ", pp.height, 
	"\nAncho: ", pp.width, 
	"\nLargo: ", pp.length, 
	"\nVolumen: ", pp.GetVolume())
}

