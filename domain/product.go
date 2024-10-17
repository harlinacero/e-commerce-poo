package domain

type Product struct {
	code           string  // Codigo
	name           string  // Nombre
	description    string  // Descripcion
	grossValue     float64 // Valor Bruto
	discPercentage float64 // Porcentage de descuento
	taxPercentage  float64 // Porcentage de impuesto
	stock          int     // Stock
}

// Constructor
func NewProduct(code string, name string, description string,
	grossvalue float64, discPercentage float64, taxPercentage float64, stock int,
) *Product {
	return &Product{
		code:           code,
		name:           name,
		description:    description,
		grossValue:     grossvalue,
		discPercentage: discPercentage,
		taxPercentage:  taxPercentage,
		stock:          stock,
	}
}

// Setters
func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) SetDescription(description string) {
	p.description = description
}

func (p *Product) SetGrossValue(value float64) {
	p.grossValue = value
}

func (p *Product) SetDisccPercentage(discPercentage float64) {
	p.discPercentage = discPercentage
}

func (p *Product) SetTaxPercentage(taxPercentage float64) {
	p.taxPercentage = taxPercentage
}

// Getters
func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetDescription() string {
	return p.description
}

func (p *Product) GetGrossValue() float64 {
	return p.grossValue
}

// Calcula el valor del descuento
func (p *Product) GetDiscValue() float64 {
	return (p.grossValue * p.discPercentage) / 100
}

// Calcula el valor del impuesto
func (p *Product) GetTaxAddValue() float64 {
	return (p.grossValue * p.taxPercentage) / 100
}

// Calcula el valor real del producto
func (p *Product) GetRealValue() float64 {
	return p.grossValue + p.GetTaxAddValue() - p.GetDiscValue()
}
