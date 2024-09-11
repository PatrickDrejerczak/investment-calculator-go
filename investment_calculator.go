package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	const inflationRate = 2.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value is", futureValue)
	fmt.Println("Future real value considering an inflation rate of 2.5% is", futureRealValue)
}
