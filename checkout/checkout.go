package checkout

import "github.com/nickWoott/go-supermarket-checkout-kata/pricing"

type Checkout struct {
	Items        map[string]int
	PricingRules pricing.PricingRules
}

func NewCheckout(pricingRules pricing.PricingRules) *Checkout {
	return &Checkout{
		Items:        make(map[string]int),
		PricingRules: pricingRules,
	}
}

func (c *Checkout) Scan(SKU string) error {
	c.Items[SKU]++
	return nil
}
