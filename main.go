package main

import "fmt"

func main() {
}

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) {
	if successRate > 0 {
		fmt.Println(float64(productionRate) * (successRate / 100))
	}
	panic("CalculateWorkingCarsPerHour not implemented")
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) {
	if successRate > 0 {
		fmt.Println((productionRate / 60) * int(successRate))
	}
	panic("CalculateWorkingCarsPerMinute not implemented")
}
