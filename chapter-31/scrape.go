package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Visited struct {
	// mu guards the visited map.
	mu sync.Mutex
	visited map[string]int
}

// Visitlink tracks that the page with the given URL has
// been visited, and returns the updated link count.

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()

	fmt.Printf("Visiting %s ...\n", url)
	
	time.Sleep(time.Duration(rand.Intn(10000) * int(time.Millisecond)))
	count := v.visited[url]
	count++
	v.visited[url] = count
	fmt.Printf("Finished visiting %s ...\n", url)
	fmt.Printf("Visited %s %d\n", url, count)
	return count
}

func main() {
	v := Visited{visited: make(map[string]int)}
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")
	go v.VisitLink("https://google.com")

	time.Sleep(1 * time.Minute)

}