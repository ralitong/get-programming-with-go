package main

import "fmt"


func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
	
	// fmt.Println(count) -- will cause an error since count is out-of-scope

}