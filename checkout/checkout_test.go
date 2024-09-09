package checkout_test

import (
	"testing"

	"github.com/nickWoott/go-supermarket-checkout-kata/checkout"
	"github.com/nickWoott/go-supermarket-checkout-kata/pricing"
)

func TestCheckout(t *testing.T) {
	t.Run("instantiates Checkout struct", func(t *testing.T) {
		co := checkout.NewCheckout(expectedRules)

		if co == nil {
			t.Errorf("Expected a Checkout struct, but got nil")
		}

		t.Run("Checkout contains correct fields", func(t *testing.T) {
			expectedRules := pricing.PricingRules{
				"A": {UnitPrice: 50, SpecialPriceQuantity: 3, SpecialPriceAmount: 130},
				"B": {UnitPrice: 30, SpecialPriceQuantity: 2, SpecialPriceAmount: 45},
				"C": {UnitPrice: 20},
				"D": {UnitPrice: 15},
			}

			co := checkout.NewCheckout(expectedRules)

			if co.items == nil {
				t.Errorf("items is nil")
			}

			if len(co.items) != 0 {
				t.Errorf("Expected empty items map, but got %v", len(co.items))
			}

			if co.pricingRules == nil {
				t.Errorf("pricingRules is nil")
			}

			_, ok := co.pricingRules.(pricing.PricingRules)
			if !ok {
				t.Errorf("pricingRules is not of type PricingRules")
			}
		})
	})
}
