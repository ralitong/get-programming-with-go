package main

import (
	"fmt"
)

func main() {
	plainText := "your message goes here"
	keyword := "golang"
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(plainText); i++ {
		currentChar := plainText[i]
		keyword := keyword[i % len(keyword)]
		shift := int(keyword - byte('a'))
		currentCharIndex := 0
		for i := 0; i < len(alphabet); i++ {
			if(currentChar == alphabet[i]) {
				currentCharIndex = i
				break
			}
		}

		var shifted byte
		if(currentChar != ' ') {
			shifted = alphabet[(currentCharIndex + shift) % len(alphabet)]
		} else {
			shifted = ' '
		}

		fmt.Printf("%c %c %c\n", currentChar, keyword, shifted)


	}
	



}