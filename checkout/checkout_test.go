package checkout_test

import (
	"fmt"
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
		expectError bool
	}{
		{"A", 1, false},
		{"B", 1, false},
		{"C", 1, false},
		{"E", 0, true},
	}

	for _, test := range tests {
		t.Run("Scan", func(t *testing.T) {
			err := co.Scan(test.sku)

			if (err != nil) != test.expectError {
				t.Errorf("Scan(%q) returned error %v, expected error: %v", test.sku, err, test.expectError)
			}

			if !test.expectError && co.Items[test.sku] != test.expectedQty {
				t.Errorf("Scan(%q) resulted in quantity %d, expected %d", test.sku, co.Items[test.sku], test.expectedQty)
			}
		})
	}
}

func TestCheckout_GetTotalPrice(t *testing.T) {
	expectedRules := pricing.NewPricingRules()
	co := checkout.NewCheckout(expectedRules)

	tests := []struct {
		scans         []string
		expectedTotal int
	}{
		{
			scans:         []string{"A"},
			expectedTotal: 50,
		},
		{
			scans:         []string{"A", "B"},
			expectedTotal: 80,
		},
		{
			scans:         []string{"A", "A", "A"},
			expectedTotal: 130,
		},
		{
			scans:         []string{"B", "B"},
			expectedTotal: 45,
		},
		{
			scans:         []string{"A", "B", "C"},
			expectedTotal: 50 + 30 + 20,
		},
		{
			scans:         []string{"A", "A", "B", "C"},
			expectedTotal: 130 + 30 + 20,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("TotalPrice with %v", test.scans), func(t *testing.T) {
			for _, sku := range test.scans {
				co.Scan(sku)
			}

			total := co.GetTotalPrice()

			if total != test.expectedTotal {
				t.Errorf("result = %d, expected %d", total, test.expectedTotal)
			}
		})
	}
}
