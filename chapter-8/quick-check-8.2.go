package main

import (
	"fmt"
	"math/big"
)

func main() {
	firstWay := big.NewInt(86400)
	fmt.Println(firstWay)

	secondWay := new(big.Int)
	secondWay.SetString("86400", 10)
	fmt.Println(secondWay)
}
