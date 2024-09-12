package checkout

import (
	"errors"
)

type PricingService interface {
	ApplyPricingRule(sku string, quantity int) (int, error)
	IsValidSKU(sku string) bool
}

type Checkout struct {
	items          map[string]int
	pricingService PricingService
}

func NewCheckout(pricingService PricingService) *Checkout {
	return &Checkout{
		items:          make(map[string]int),
		pricingService: pricingService,
	}
}

func (c *Checkout) Scan(SKU string) error {
	isValid := c.pricingService.IsValidSKU(SKU)

	if !isValid {
		return errors.New("invalid sku")
	}
	c.items[SKU]++
	return nil
}

func (c *Checkout) GetTotalPrice() (int, error) {

	if len(c.items) <= 0 {
		return 0, errors.New("no items scanned")
	}

	totalPrice := 0

	for SKU, quantity := range c.items {
		finalItemPrice, err := c.pricingService.ApplyPricingRule(SKU, quantity)

		if err != nil {
			return 0, err
		}

		totalPrice += finalItemPrice
	}

	return totalPrice, nil
}
