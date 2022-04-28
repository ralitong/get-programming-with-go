package main

import "fmt"

func main() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that")
	}

	var adventureRoom = "mountain"

	if adventureRoom == "cave" {
		fmt.Println("Caves are damp dark places")
	} else if adventureRoom == "entrance" {
		fmt.Println("You are near the entrance")
	} else if adventureRoom == "mountain" {
		fmt.Println("Mountains are high places with low oxygen")
	} else {
		fmt.Println("The place", adventureRoom, "does not exist")
	}
}
