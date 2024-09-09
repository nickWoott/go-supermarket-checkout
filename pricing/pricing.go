package pricing

type Rule struct {
	UnitPrice            int
	SpecialPriceQuantity int
	SpecialPriceAmount   int
}

type PricingRules map[string]Rule

func NewPricingRules() PricingRules {
	return PricingRules{
		"A": {UnitPrice: 50, SpecialPriceQuantity: 3, SpecialPriceAmount: 130},
		"B": {UnitPrice: 30, SpecialPriceQuantity: 2, SpecialPriceAmount: 45},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}
}
