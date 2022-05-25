package main

import (
	"fmt"
	"math"
)

type report struct {
	sol
	temperature
	location
}

type sol int

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}

	return days
}

type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type location struct {
	lat, long float64
}

func (l location) days(l2 location) int {
	s1, c1 := math.Sincos(rad(l.lat))
	s2, c2 := math.Sincos(rad(l2.lat))
	clong := math.Cos(rad(l.long - l2.long))
	return int(3389.5 * math.Acos(s1*s2+c1*c2*clong))
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type celsius float64

func main() {
	report := report{sol:15}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
}