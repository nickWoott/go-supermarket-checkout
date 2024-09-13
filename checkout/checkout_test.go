package checkout_test

import (
	"testing"

	"github.com/nickWoott/go-supermarket-checkout-kata/checkout"
	"github.com/nickWoott/go-supermarket-checkout-kata/pricing"
)

func TestCheckout_GetTotalPrice(t *testing.T) {

	tests := []struct {
		name          string
		scans         []string
		expectedTotal int
		expectedError bool
	}{
		{
			name:          "Single item scan",
			scans:         []string{"A"},
			expectedTotal: 50,
			expectedError: false,
		},
		{
			name:          "Multiple items scan",
			scans:         []string{"A", "B"},
			expectedTotal: 80,
			expectedError: false,
		},
		{
			name:          "Special price with multiple same items",
			scans:         []string{"A", "A", "A"},
			expectedTotal: 130,
			expectedError: false,
		},
		{
			name:          "Multiple B items",
			scans:         []string{"B", "B"},
			expectedTotal: 45,
			expectedError: false,
		},
		{
			name:          "Mixed items scan",
			scans:         []string{"A", "B", "C"},
			expectedTotal: 50 + 30 + 20,
			expectedError: false,
		},
		{
			name:          "Mixed items scan with duplicate",
			scans:         []string{"A", "A", "B", "C"},
			expectedTotal: 50 + 50 + 30 + 20,
			expectedError: false,
		},
		{
			name:          "No items scanned",
			scans:         []string{},
			expectedTotal: 0,
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ps := pricing.NewPricingService()
			co := checkout.NewCheckout(ps)
			for _, sku := range test.scans {
				co.Scan(sku)
			}

			total, err := co.GetTotalPrice()

			if (err != nil) != test.expectedError {
				t.Errorf("GetTotalPrice() error = %v, expectedError %v", err, test.expectedError)
			}

			if total != test.expectedTotal {
				t.Errorf("result = %d, expected %d", total, test.expectedTotal)
			}
		})
	}
}
