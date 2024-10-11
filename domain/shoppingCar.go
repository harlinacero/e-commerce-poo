package domain

type ShoppingCar struct {
	Productcs []Product
}

// NewShoppingCar returns new ShoppoingCar.
func NewShoppingCar(produts []Product) *ShoppingCar {
	return &ShoppingCar{
		Productcs: produts,
	}
}

// Add product to shopping car
func (sc *ShoppingCar) AddProduct(product Product) {
	sc.Productcs = append(sc.Productcs, product)
}

// Remove product of shopping car
func (sc *ShoppingCar) RemoveProduct(productName string) {
	for i, product := range sc.Productcs {
		if product.Name == productName {
			sc.Productcs = append(sc.Productcs[:i], sc.Productcs[i+1:]...)
		}
	}
}

func (sc *ShoppingCar) CalculateTotalValue() float64 {
	var total float64
	for _, product := range sc.Productcs {
		total += product.GetRealValue()
	}

	return total
}
