package main

import (
	"fmt"
	"math"
)

type rover struct {
	gps
}

func (r rover) status() {
	fmt.Print(r.gps.message())
}

type gps struct {
	current location
	destination location
	world
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.2f km remaining from %s to %s\n",g.distance(), g.current.name, g.destination.name)
}

type world struct {
	radius float64
}


type location struct {
	name string
	lat, long float64
}

func (w world) distance(l1, l2 location) float64 {
	s1, c1 := math.Sincos(rad(l1.lat))
	s2, c2 := math.Sincos(rad(l2.lat))
	clong := math.Cos(rad(l1.long - l2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (l location) description() string {
	return fmt.Sprintf("name: %s latitude: %.2f longitude %.2f\n", l.name, l.lat, l.long)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}


func main() {
	mars := world{radius: 3389.5}
	bradbury := location{"Bradbury Landing", -4.5895, 137.4417}
	planitia := location{"Elysium Planitia", 4.5, 135.9}
	gps := gps{bradbury, planitia, mars}
	curiosity := rover{gps}
	curiosity.status()
}