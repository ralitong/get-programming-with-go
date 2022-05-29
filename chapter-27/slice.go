package main

import "fmt"

func main() {
	var soup []string
	fmt.Println("The slice is nil:",soup == nil)

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println("The length of soup is:",len(soup))
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)
}