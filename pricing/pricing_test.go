package pricing_test

import (
	"testing"

	"github.com/nickWoott/go-supermarket-checkout-kata/pricing"
)

func TestPricing(t *testing.T) {
	t.Run("return empty PricingRules struct", func(t *testing.T) {
		pr := pricing.NewPricingRules()

		if pr == nil {
			t.Fatalf("Expected PricingRules struct, got nil")
		}
	})
}
