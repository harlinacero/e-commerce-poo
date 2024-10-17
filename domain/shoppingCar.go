package domain

type ShoppingCar struct {
	user	  User
	productcs []Product
}

// NewShoppingCar returns new ShoppoingCar.
func NewShoppingCar(user User, produts []Product) *ShoppingCar {
	return &ShoppingCar{
		user: user,
		productcs: produts,
	}
}

// Add product to shopping car
func (sc *ShoppingCar) AddProduct(product Product) {
	sc.productcs = append(sc.productcs, product)
}

// Remove product of shopping car
func (sc *ShoppingCar) RemoveProduct(productName string) {
	for i, product := range sc.productcs {
		if product.GetName() == productName {
			sc.productcs = append(sc.productcs[:i], sc.productcs[i+1:]...)
		}
	}
}

func (sc *ShoppingCar) GetProducts() []Product {
	return sc.productcs
}

func (sc *ShoppingCar) CalculateTotalValue() float64 {
	var total float64
	for _, product := range sc.productcs {
		total += product.GetRealValue()
	}

	return total
}
