package main

import (
	"fmt"
	"strings"
)


func sourceGopher(downstream chan string) {
	for _, v := range []string{
		"hello world",
		"a bad apple", 
		"goodbye all", "a", "b","a", "c"} {
		downstream <- v
	}

	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	memory := make(map[string]bool)

	for item := range upstream {
		if !strings.Contains(item, "bad") && !memory[item] {
			memory[item] = true
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {

	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)

}