package checkout

import (
	"errors"

	"github.com/nickWoott/go-supermarket-checkout-kata/pricing"
)

type ICheckout interface {
	Scan(SKU string) error
	GetTotalPrice() (totalPrice int, err error)
}

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
	if _, exists := c.PricingRules[SKU]; !exists {
		return errors.New("invalid item ")
	}
	c.Items[SKU]++
	return nil
}

func (c *Checkout) GetTotalPrice() (int, error) {

	if len(c.Items) <= 0 {

		return 0, errors.New("no items scanned")

	}

	totalPrice := 0

	for SKU, count := range c.Items {
		rule := c.PricingRules[SKU]

		if rule.SpecialPriceQuantity > 0 {
			specialPrice := (count / rule.SpecialPriceQuantity) * rule.SpecialPriceAmount
			standardPrice := (count % rule.SpecialPriceQuantity) * rule.UnitPrice

			totalPrice += specialPrice + standardPrice
		} else {
			totalPrice += count * rule.UnitPrice
		}
	}

	return totalPrice, nil
}
