package main

import (
	"fmt"
	"math"
)

func main() {
	bh := float64(32767)
	h := int16(bh)

	fmt.Println(h)

	if bh < math.MinInt16 && bh > math.MaxInt16 {
		fmt.Println("bh will overflow if converted to int16")
	} else {
		fmt.Println("bh fits if converted to int16")
	}
}
