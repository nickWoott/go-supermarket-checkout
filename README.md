# About

This is a simple program replicating a checkout system. The program is run in the command line, and is able to scan SKU codes of items, eventually returning the total price. 


## Running the program

To install dependencies:

`go mod tidy`

To run the CLI implementation, from the root of the project run:

`go run ./main.go`

The program will allow entry of any SKU found in the current pricing rules. Upon typing 'done', the GetTotalPrice method will be called. 

## Tests

Tests can be found for both the pricing and checkout packages.

`go test ./checkout/checkout_test.go`
`go test ./pricing/pricing_test.go`

## Structure and developing in this repo

The repository consists of two main packages, checkout and pricing. Pricing holds the pricing rules, and checkout contains the checkout functionality to create the checkout struct along with the methods for scanning and returning the total price. 

The main package contains the CLI implementation.
