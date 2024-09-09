package checkout_test

import (
	"testing"

	"github.com/nickWoott/go-supermarket-checkout-kata/checkout"
	"github.com/nickWoott/go-supermarket-checkout-kata/pricing"
)

func TestCheckout_NewCheckout(t *testing.T) {
	expectedRules := pricing.PricingRules{
		"A": {UnitPrice: 50, SpecialPriceQuantity: 3, SpecialPriceAmount: 130},
		"B": {UnitPrice: 30, SpecialPriceQuantity: 2, SpecialPriceAmount: 45},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}
	t.Run("instantiates Checkout struct", func(t *testing.T) {
		co := checkout.NewCheckout(expectedRules)

		if co == nil {
			t.Errorf("Expected a Checkout struct, but got nil")
		}

		t.Run("Checkout contains correct fields", func(t *testing.T) {

			co := checkout.NewCheckout(expectedRules)

			if co.Items == nil {
				t.Errorf("items is nil")
			}

			if len(co.Items) != 0 {
				t.Errorf("Expected empty items map, but got %v", len(co.Items))
			}

			if co.PricingRules == nil {
				t.Errorf("pricingRules is nil")
			}

			if _, ok := interface{}(co.PricingRules).(pricing.PricingRules); !ok {
				t.Errorf("pricingRules is not of type pricing.PricingRules")
			}
		})
	})
}

func TestCheckout_Scan(t *testing.T) {
	expectedRules := pricing.NewPricingRules()
	co := checkout.NewCheckout(expectedRules)

	tests := []struct {
		sku         string
		expectedQty int
	}{
		{"A", 1},
		{"B", 1},
		{"C", 1},
	}

	for _, test := range tests {
		t.Run("Scan appends string to items field", func(t *testing.T) {
			err := co.Scan(test.sku)

			if err != nil {
				t.Errorf("Scan(%q) returned error %v, expected no error", test.sku, err)
			}

			if qty := co.Items[test.sku]; qty != test.expectedQty {
				t.Errorf("Scan(%q) resulted in quantity %d, expected %d", test.sku, qty, test.expectedQty)
			}
		})
	}
}
