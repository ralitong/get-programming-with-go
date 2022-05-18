package main

import "fmt"


func main() {
	temperature := map[string]int{
		"Earth" : 15,
		"Mars" : -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v degrees C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	// a non-existing key in the map
	moon := temperature["Moon"]
	fmt.Println(moon)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v degrees C.\n", moon)
	} else {
		fmt.Println("Where is the moon")
	}
	
	temperature["Moon"] = 10
	
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v degrees C.\n", moon)
	} else {
		fmt.Println("Where is the moon")
	}
}