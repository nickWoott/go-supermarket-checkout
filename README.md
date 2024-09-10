# About

This is a simple program replicating a checkout system. The program is run in the command line, and is able to scan SKU codes of items, eventually returning the total price. 

## Structure 

The repository consists of two main packages, checkout and pricing. Pricing holds the pricing rules, and checkout contains the checkout functionality to create the checkout struct along with the methods for scanning and returning the total price. 

The main package contains the CLI implementation.

## Running the program

To install dependencies:

`go mod tidy`

To run the CLI implementation, from the root of the project run:

`go run ./main.go`

The program will allow entry of any SKU found in the current pricing rules. Upon typing 'done', the GetTotalPrice method will be called. 

Note that SKU's are case sensitive. 

## Tests

Tests can be found for both the pricing and checkout packages.

`go test ./checkout/checkout_test.go`
`go test ./pricing/pricing_test.go`

## Future Improvements

This kata was completed using TDD with functionality in mind. These tests rely on concrete implementation of the Checkout and PricingRules structs. To make this a more extensible and flexible program, the next step would be to add interface implementation for the Checkout and PricingRules, then amending the testing suite to only test method implementation. 


