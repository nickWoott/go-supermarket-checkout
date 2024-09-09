package checkout_test

import (
	"testing"

	"github.com/nickWoott/go-supermarket-checkout-kata/checkout"
)

func TestCheckout(t *testing.T) {
	t.Run("return empty Checkout struct", func(t *testing.T) {
		co := checkout.NewCheckout()

		if co == nil {
			t.Fatalf("Expected a Checkout struct, but got nil")
		}

	})
}
