package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturn float64
	var years float64

	fmt.Print("Put amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturn)
	fmt.Print("For how many years: ")
	fmt.Scan(&years)

	result := investmentAmount * math.Pow((1+expectedReturn/100), years)
	futureRealValue := result / math.Pow(1+inflationRate/100, years)

	fmt.Println("Result", result)
	fmt.Println("Real value", futureRealValue)
}
