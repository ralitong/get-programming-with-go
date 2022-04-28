package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {

	for count:= 10; count > 0; count-- {
		year := rand.Intn(1000) + 2000
		isLeap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2:
			if isLeap {
				fmt.Println("The year is leap")
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 7, 9, 11:
			daysInMonth = 30
		}
	
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}



}