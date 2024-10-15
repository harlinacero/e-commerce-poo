package domain

type Product struct {
	code        string
	name        string
	description string
	grossValue  float64
	discValue   float64
	taxAddValue float64
	stock       int
}

func NewProduct(code string, name string, description string,
	grossvalue float64, discValue float64, taxAddValue float64, stock int) *Product {
	return &Product{
		code:        code,
		name:        name,
		description: description,
		grossValue:  grossvalue,
		discValue:   discValue,
		taxAddValue: taxAddValue,
		stock:       stock,
	}
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) SetDescription(description string) {
	p.description = description
}

func (p *Product) SetGrossValue(value float64) {
	p.grossValue = value
}

func (p *Product) SetDiscValue(discValue float64) {
	p.discValue = discValue
}

func (p *Product) SetTaxAddValue(taxAddValue float64) {
	p.taxAddValue = taxAddValue
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetDescription() string {
	return p.description
}

func (p *Product) GetGrossValue() float64 {
	return p.grossValue
}

func (p *Product) GetDiscValue() float64 {
	return p.discValue * p.grossValue
}

func (p *Product) GetTaxAddValue() float64 {
	return p.taxAddValue * p.taxAddValue
}

func (p *Product) GetRealValue() float64 {
	return p.grossValue + p.GetTaxAddValue() - p.GetDiscValue()
}
