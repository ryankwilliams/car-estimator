// Package provides functions used to calculate a cars total cost to purchase
package carestimator

import "fmt"

// A car estimator struct to store all of the required data fields needed
// to calculate the total cost for a new car
type carEstimator struct {
	carPrice      float64
	tradePrice    float64
	stateSalesTax float64
	fees          float64
	downPayment   float64
}

// Gets the purchasing car's price from the user and returns it
func getCarPrice() float64 {
	var carPrice float64
	fmt.Print("Enter the car cost: ")
	fmt.Scan(&carPrice)
	return carPrice
}

// Gets the trade in car's price from the user and returns it
func getTradePrice() float64 {
	var tradePrice float64
	fmt.Print("Enter the trade in price: ")
	fmt.Scan(&tradePrice)
	return tradePrice
}

// Gets the state sales tax and returns it
func getStateSalesTax() float64 {
	var salesTax float64
	fmt.Print("Enter the state sales tax: ")
	fmt.Scan(&salesTax)
	return salesTax
}

// Gets the total fee estimate and returns it
func getFeeEstimates() float64 {
	var fees float64
	fmt.Print("Enter the total fees (e.g. title, doc fees, etc): ")
	fmt.Scan(&fees)
	return fees
}

// Gets the down payment and returns it
func getDownPayment() float64 {
	var downPayment float64
	fmt.Print("Enter the down payment: ")
	fmt.Scan(&downPayment)
	return downPayment
}

// Gets all required input from the user and returns the car estimator struct
func getInput(car carEstimator) carEstimator {
	fmt.Println("** Car Estimator **")

	car.carPrice = getCarPrice()
	car.tradePrice = getTradePrice()
	car.stateSalesTax = getStateSalesTax()
	car.fees = getFeeEstimates()
	car.downPayment = getDownPayment()
	return car
}

// Calculates the total cost of the car to be purchased
func CalculateTotalCost() {
	var car carEstimator
	var totalCost float64
	var carCostWithTradeTotal float64
	var salesTaxTotal float64

	car = getInput(car)

	carCostWithTradeTotal = car.carPrice - car.tradePrice
	salesTaxTotal = carCostWithTradeTotal * car.stateSalesTax
	totalCost = carCostWithTradeTotal + salesTaxTotal + car.fees - car.downPayment

	fmt.Print("\n--------------------------------------------------\n",
	"New Car Cost                    : $", car.carPrice, "\n",
	"Current Car Trade In Value      : $", car.tradePrice, "\n",
	"New Car Total Cost (w/trade)    : $", carCostWithTradeTotal, "\n",
	"Total Sales Tax (new car total) : $", salesTaxTotal, "\n",
	"Total purchasing fees           : $", car.fees, "\n",
	"Total down payment              : $", car.downPayment, "\n",
	"--------------------------------------------------\n",
	"Total Purchase Cost             : $", totalCost, "\n",
	)
}
