package main

import (
	"fmt"
	"math/rand"
)


func main() {

	const nickel = 5
	const dime = 10
	const quarter = 25

	coins := []int{nickel, dime, quarter}
	piggyBank := 0

	fmt.Printf("%-7v %15v \n", "Deposit", "Running Balance")
	fmt.Println("======================")
	
	for {
		randomCoin := coins[rand.Intn(3)]
		toDollar := piggyBank / 100
		remainder := piggyBank % 100
		fmt.Printf("%2v¢ %17v.%v$ \n", randomCoin, toDollar, remainder)
		if((piggyBank / 100) >= 20.0) {
			break;
		}
		
		piggyBank += randomCoin
	}
}