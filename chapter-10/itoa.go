package main

import (
	"fmt"
	"strconv"
)


func main() {
	countdown := 10

	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Print(str + "\n")

	str = fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str)

	otherway, err := strconv.Atoi("10")
	if err != nil {
		// oh no, something went wrong
	}
	fmt.Println(otherway)
}