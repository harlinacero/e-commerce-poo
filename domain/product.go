package domain

type Product struct {
	Code        string
	Name        string
	Description string
	GrossValue  float64
	DiscValue   float64
	TaxAddValue float64
	Stock       int
}

func NewProduct(code string, name string, description string,
	grossvalue float64, discValue float64, taxAddValue float64, stock int) *Product {
	return &Product{
		Code:        code,
		Name:        name,
		Description: description,
		GrossValue:  grossvalue,
		DiscValue:   discValue,
		TaxAddValue: taxAddValue,
		Stock:       stock,
	}
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) SetDescription(description string) {
	p.Description = description
}

func (p *Product) SetGrossValue(value float64) {
	p.GrossValue = value
}

func (p *Product) SetDiscValue(discValue float64) {
	p.DiscValue = discValue
}

func (p *Product) SetTaxAddValue(taxAddValue float64) {
	p.TaxAddValue = taxAddValue
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) GetGrossValue() float64 {
	return p.GrossValue
}

func (p *Product) GetDiscValue() float64 {
	return p.DiscValue * p.GrossValue
}

func (p *Product) GetTaxAddValue() float64 {
	return p.TaxAddValue * p.TaxAddValue
}

func (p *Product) GetRealValue() float64 {
	return p.GrossValue + p.GetTaxAddValue() - p.GetDiscValue()
}
