package product



type ProductBase struct {
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
) *ProductBase {
	return &ProductBase{
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
func (p *ProductBase) SetName(name string) {
	p.name = name
}

func (p *ProductBase) SetDescription(description string) {
	p.description = description
}

func (p *ProductBase) SetGrossValue(value float64) {
	p.grossValue = value
}

func (p *ProductBase) SetDisccountPercentage(discPercentage float64) {
	p.discPercentage = discPercentage
}

func (p *ProductBase) SetTaxPercentage(taxPercentage float64) {
	p.taxPercentage = taxPercentage
}

func (p *ProductBase) SetStock(stock int) {
	p.stock = stock
}
// Getters
func (p *ProductBase) GetName() string {
	return p.name
}

func (p *ProductBase) GetDescription() string {
	return p.description
}

func (p *ProductBase) GetGrossValue() float64 {
	return p.grossValue
}

// Calcula el valor del descuento
func (p *ProductBase) GetDiscValue() float64 {
	return (p.grossValue * p.discPercentage) / 100
}

// Calcula el valor del impuesto
func (p *ProductBase) GetTaxAddValue() float64 {
	return (p.grossValue * p.taxPercentage) / 100
}

// Calcula el valor real del Producto
func (p *ProductBase) GetRealValue() float64 {
	return p.grossValue + p.GetTaxAddValue() - p.GetDiscValue()
}

func (p *ProductBase) GetStock() int {
	return p.stock
}