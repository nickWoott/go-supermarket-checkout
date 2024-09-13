package pricing_test

import (
	"testing"

	"github.com/nickWoott/go-supermarket-checkout-kata/pricing"
)

func TestPricing_ApplyPricingRule(t *testing.T) {

	tests := []struct {
		sku         string
		quantity    int
		expected    int
		expectError bool
	}{

		{"A", 1, 50, false},
		{"A", 3, 130, false},
		{"A", 4, 180, false},

		{"B", 1, 30, false},
		{"B", 2, 45, false},
		{"B", 3, 75, false},

		{"C", 1, 20, false},
		{"C", 2, 40, false},

		{"E", 1, 0, true},
	}

	ps := pricing.NewPricingService()

	for _, test := range tests {
		t.Run(test.sku, func(t *testing.T) {
			price, err := ps.ApplyPricingRule(test.sku, test.quantity)

			if (err != nil) != test.expectError {
				t.Errorf("ApplyPricingRule(%q, %d) returned error %v, expected error: %v", test.sku, test.quantity, err, test.expectError)
			}

			if price != test.expected {
				t.Errorf("ApplyPricingRule(%q, %d) = %d, expected %d", test.sku, test.quantity, price, test.expected)
			}
		})
	}
}

func TestPricing_IsValidSku(t *testing.T) {

	tests := []struct {
		sku      string
		expected bool
	}{
		{"A", true}, {"B", true}, {"C", true}, {"D", true}, {"E", false},
	}

	ps := pricing.NewPricingService()

	for _, test := range tests {
		t.Run(test.sku, func(t *testing.T) {

			isValidSKU := ps.IsValidSKU(test.sku)

			if isValidSKU != test.expected {
				t.Errorf("IsValidSKU returned %v, expected %v, when passed %s", isValidSKU, test.expected, test.sku)
			}

		})
	}
}
