package main

import (
	"fmt"
	"log"

	"github.com/nickWoott/go-supermarket-checkout-kata/checkout"
	"github.com/nickWoott/go-supermarket-checkout-kata/pricing"
)

func main() {

	type ICheckout interface {
		Scan(SKU string) error
		GetTotalPrice() (totalPrice int, err error)
	}

	ps := pricing.NewPricingService()
	var co ICheckout = checkout.NewCheckout(ps)

	fmt.Println("Enter items to scan. Type 'done' to finish:")

	for {
		var input string
		fmt.Print("Enter item SKU: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		if input == "done" {
			break
		}

		err = co.Scan(input)
		if err != nil {
			fmt.Printf("Error scanning item %s: %v\n", input, err)
		}
	}

	total, err := co.GetTotalPrice()
	if err != nil {
		log.Fatalf("Error getting total price: %v", err)
	}

	fmt.Printf("Total price: %d\n", total)
}
