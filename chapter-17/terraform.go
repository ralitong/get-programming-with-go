package main

import "fmt"

type Planets []string

func (planets Planets) Terraform() {
	for i, planet := range planets {
		planets[i] = "New " + planet
	}
}

func main() {
	planets := Planets{"Mars", "Uranus", "Neptune"}
	planets.Terraform()
	fmt.Println(planets)

}