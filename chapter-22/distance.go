package main

import (
	"fmt"
	"math"
	"sort"
)

type location struct {
	name      string
	lat, long float64
}

func newLocation(lander string, lat, long coordinate) location {
	return location{lander, lat.decimal(), long.decimal()}
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}


type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}


func main() {

	worlds := map[string]world{
		"mercury": {2439.7},
		"venus":{6051.8},
		"earth":{6371.0},
		"mars":{3389.5},
		"jupiter":{69911},
		"uranus":{25362},
		"neptune":{24622},
	}

	locations := []location{
		{"Columbia Memorial Station", coordinate{14, 34, 6.2, 'S'}.decimal(), coordinate{175, 28, 21.5, 'E'}.decimal()},
		{"Challenger Memorial Station", coordinate{1, 56, 46.3, 'S'}.decimal(), coordinate{354, 28, 24.2, 'E'}.decimal()},
		{"Bradbury Landing", coordinate{4, 35, 22.2, 'S'}.decimal(), coordinate{137, 26, 30.1, 'E'}.decimal()},
		{"Elysium Planitia", coordinate{4, 30, 0.2, 'N'}.decimal(), coordinate{135, 54, 0, 'E'}.decimal()},
	}

	combis := Combinations(len(locations), 2)
	marsDistances := make(map[string]float64, len(combis))


	for i := 0; i < len(combis); i++ {
		combi := combis[i]
		firstLocation := locations[combi[0]]
		secondLocation := locations[combi[1]]
		distance := worlds["mars"].distance(firstLocation, secondLocation)
		
		label := firstLocation.name + " to " + secondLocation.name
		marsDistances[label] = distance
	}


	labels := make([]string,0, len(marsDistances))
	
	for label := range marsDistances {
		labels = append(labels, label)
	}

	sort.SliceStable(labels, func(i, j int) bool {
		return marsDistances[labels[i]] < marsDistances[labels[j]]
	})

	for _, label := range labels {
		fmt.Printf("%s: %.2f km\n", label, marsDistances[label])
	}

	england := location{"England", coordinate{51, 30, 0, 'N'}.decimal(), coordinate{0, 8, 0, 'W'}.decimal()}
	france := location{"France", coordinate{48, 51, 0, 'N'}.decimal(), coordinate{2, 21, 0, 'E'}.decimal()}
	fmt.Printf("Distance between %s and %s is %.2f km\n", england.name, france.name, worlds["earth"].distance(england, france))
	
	
	manila := location{"Manila", coordinate{14, 35, 58.2432, 'N'}.decimal(), coordinate{120, 59, 3.1992, 'E'}.decimal()}
	helsinki := location{"Helsinki", coordinate{60, 11, 31.4124, 'N'}.decimal(), coordinate{24, 56, 44.9916, 'E'}.decimal()}
	fmt.Printf("Distance between %s and %s is %.2f km\n", manila.name, helsinki.name, worlds["earth"].distance(manila, helsinki))
	
	
	mountSharp := location{"Mount Sharp", coordinate{5, 4, 48, 'S'}.decimal(), coordinate{137, 51, 0, 'E'}.decimal()}
	olympusMons := location{"Olympus Mons", coordinate{18, 39, 0, 'N'}.decimal(), coordinate{226, 12, 0, 'E'}.decimal()}
	fmt.Printf("Distance between %s and %s is %.2f km\n", mountSharp.name, olympusMons.name, worlds["mars"].distance(mountSharp, olympusMons))



}

// Combinations generates all of the combinations of k elements from a
// set of size n. The returned slice has length Binomial(n,k) and each inner slice
// has length k.
// 
// n and k must be non-negative with n >= k, otherwise Combinations will panic.
// 
func Combinations(n, k int)[][]int {
	combins := Binomial(n, k)
	data := make([][]int, combins)
	if len(data) == 0 {
		return data
	}
	data[0] = make([]int, k)
	for i := range data[0] {
		data[0][i] = i
	}
	for i := 1; i < combins; i++ {
		next := make([]int, k)
		copy(next, data[i-1])
		nextCombination(next, n, k)
		data[i] = next
	}
	return data
}

// nextCombination generates the combination after s, overwriting the input value
func nextCombination(s []int, n, k int) {
	for j := k - 1; j >= 0; j-- {
		if s[j] == n+j-k {
			continue
		}
		s[j]++
		for l := j + 1; l < k; l++ {
			s[l] = s[j] + l - j
		}
		break
	}
}

// Binomial returns the binomial coefficient of (n, k), also commonly referred to
// as "n choose k".
// 
// The binomial coefficient, C(n,k), is the number of unordered combinations of
// k elements in a set that is n elements big, and is defined as
// 
// C(n,k) = n!/((n-k)!k!)
// 
// n and k must be non-negative with n >= k, otherwise Binomial will panic
// No check is made for overflow
func Binomial(n, k int) int {
	if n < 0 || k < 0 {
		panic("negative input")
	}
	if n < k {
		panic("n is less than k")
	}

	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}

	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}

	return b
}

