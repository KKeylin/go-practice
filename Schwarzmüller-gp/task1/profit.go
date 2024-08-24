package main

import (
	"fmt"
)

func main() {
	const inflationRate = 2.5
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	beforeTax := (revenue - expenses)
	afterTax := beforeTax - (beforeTax * taxRate / 100)

	fmt.Println("Result", beforeTax)
	fmt.Println("Real value", afterTax)
}
