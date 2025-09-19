package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var interestRate float64
	var years float64

	const inflationRate = 2.5

	fmt.Print("Enter investment amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter interest rate:")
	fmt.Scan(&interestRate)

	fmt.Print("Enter years:")
	fmt.Scan(&years)

	annualReturn := investmentAmount * math.Pow(1+interestRate/100, years)
	annualRealReturn := annualReturn / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Annual Return: %.2f\n", annualReturn)
	fmt.Printf("Annual Real Return: %.2f\n", annualRealReturn)
}
