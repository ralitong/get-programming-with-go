package main

import (
	"fmt"
	"math/rand"
)


func main() {

	const nickel = 0.05
	const dime = 0.10
	const quarter = 0.25

	coins := []float64{nickel, dime, quarter}
	piggyBank := 0.0

	fmt.Printf("%-7v %15v \n", "Deposit", "Running Balance")
	fmt.Println("======================")
	
	for {
		randomCoin := coins[rand.Intn(3)]
		fmt.Printf("%.2f %17.2f \n", randomCoin, piggyBank)
		if(piggyBank >= 20.0) {
			break;
		}
		
		piggyBank += randomCoin
	}
}