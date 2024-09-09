package pricing_test

import (
	"testing"

	"github.com/nickWoott/go-supermarket-checkout-kata/pricing"
)

func TestPricing(t *testing.T) {
	t.Run("returns PricingRules map", func(t *testing.T) {
		pr := pricing.NewPricingRules()

		if pr == nil {
			t.Errorf("Expected PricingRules map, got nil")
		}
	})

	t.Run("returns expected pricing rules", func(t *testing.T) {
		pr := pricing.NewPricingRules()

		expectedRules := PricingRules{
			"A": {UnitPrice: 50, SpecialPriceQuantity: 3, SpecialPriceAmount: 130},
			"B": {UnitPrice: 30, SpecialPriceQuantity: 2, SpecialPriceAmount: 45},
			"C": {UnitPrice: 20},
			"D": {UnitPrice: 15},
		}

		if len(pr) != len(expectedRules) {
			t.Errorf("Expected length %d, but got %d", len(expectedRules), len(pr))
		}

		for sku, expectedRule := range expectedRules {
			actualRule, exists := pr[sku]
			if !exists {
				t.Errorf("SKU %s not found in pricing rules", sku)
				continue
			}
			if actualRule != expectedRule {
				t.Errorf("For SKU %s, expected %+v but got %+v", sku, expectedRule, actualRule)
			}
		}
	})
}
