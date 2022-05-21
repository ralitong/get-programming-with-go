package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Printf("%v\n",opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Printf("%+v\n",insight)
}