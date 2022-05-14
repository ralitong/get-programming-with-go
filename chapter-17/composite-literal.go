package main

import "fmt"

func main() {
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	fmt.Printf("The type of dwarfArray is %T\n", dwarfArray)
	fmt.Printf("The type of dwarfSlice is %T\n", dwarfSlice)
}