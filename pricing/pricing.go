package pricing

import "fmt"

type Rule struct {
	unitPrice            int
	specialPriceQuantity int
	specialPriceAmount   int
}

type PricingRules map[string]Rule

type PricingService struct {
	pricingRules PricingRules
}

type item struct {
	sku      string
	quantity int
}

var pricingRules = PricingRules{
	"A": {unitPrice: 50, specialPriceQuantity: 3, specialPriceAmount: 130},
	"B": {unitPrice: 30, specialPriceQuantity: 2, specialPriceAmount: 45},
	"C": {unitPrice: 20},
	"D": {unitPrice: 15},
}

func NewPricingService() *PricingService {
	return &PricingService{
		pricingRules: pricingRules,
	}
}

func (p *PricingService) ApplyPricingRule(sku string, quantity int) (int, error) {

	pricingRule, exists := p.pricingRules[sku]
	if !exists {
		return 0, fmt.Errorf("pricing rule not found for SKU: %s", sku)
	}

	if pricingRule.specialPriceQuantity > 0 {

		specialPrice := (quantity / pricingRule.specialPriceQuantity) * pricingRule.specialPriceAmount

		standardPrice := (quantity % pricingRule.specialPriceQuantity) * pricingRule.unitPrice
		return specialPrice + standardPrice, nil

	} else {

		return (quantity * pricingRule.unitPrice), nil
	}

}

func (p *PricingService) IsValidSKU(sku string) bool {
	_, exists := p.pricingRules[sku]

	if !exists {
		return false
	}
	return true

}
