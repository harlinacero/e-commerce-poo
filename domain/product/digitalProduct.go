package product

import "fmt"

type DigitalProduct struct {
	ProductBase            // Hereda de producto
	fileFormat string  // Formato del archivo
	size       float64 // Tamaño del archivo
}

// Constructor
func NewDigitalProduct(code string, name string, description string, grossvalue float64, discPercentage float64, taxPercentage float64, stock int, fileFormat string, size float64) *DigitalProduct {
	return &DigitalProduct{
		ProductBase: ProductBase{
			code:           code,
			name:           name,
			description:    description,
			grossValue:     grossvalue,
			discPercentage: discPercentage,
			taxPercentage:  taxPercentage,
			stock:          stock,
		},
		fileFormat: fileFormat,
		size:       size,
	}
}

// Setters
func (dp *DigitalProduct) SetFileFormat(fileFormat string) {
	dp.fileFormat = fileFormat
}

func (dp *DigitalProduct) SetSize(size float64) {
	dp.size = size
}

// Getters
func (dp *DigitalProduct) GetFileFormat() string {
	return dp.fileFormat
}

func (dp *DigitalProduct) GetSize() float64 {
	return dp.size
}

// Mostrar información del producto
func (dp *DigitalProduct) ShowInfo() string {
	return fmt.Sprintf("Formato: %s\nTamaño: %.2f MB\n", dp.GetFileFormat(), dp.GetSize())
}
