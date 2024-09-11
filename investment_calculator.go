package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFutureValue := fmt.Sprintf("Future value is %.2f \n", futureValue)

	// fmt.Println("Future value is", futureValue)
	outputText(formattedFutureValue, futureRealValue)
}

func outputText(formattedFutureValue string, futureRealValue float64) {
	fmt.Print(formattedFutureValue)
	fmt.Printf("Future value (adjusted for inflation) %.2f \n", futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// }
