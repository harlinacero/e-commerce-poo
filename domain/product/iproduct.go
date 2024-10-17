package product

type Product interface {
	SetName(name string)
	GetName() string
	SetDescription(description string)
	SetDisccountPercentage(discPercentage float64)
	SetTaxPercentage(taxPercentage float64)
	SetGrossValue(value float64)
	SetStock(stock int)
	GetDescription() string
	GetTaxAddValue() float64
	GetDiscValue() float64
	GetGrossValue() float64
	GetRealValue() float64
	GetStock() int
	ShowInfo() string
}