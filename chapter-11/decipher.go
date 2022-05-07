package main

import "fmt"


func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	keyword := "GOLANG"

	// A B C D E F G H I J  K  L  M  N  O  P  Q  R  S  T  U  V  W  X  Y  Z
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25

	for i := 0; i < len(cipherText); i++ {
		current := cipherText[i]
		currentKeyword := keyword[i % len(keyword)]
		currentIndex := 0
		currentKeywordIndex := 0

		for i:= 0; i < len(alphabet); i++ {
			if current == alphabet[i] {
				currentIndex = i
				break
			}
		}

		for i:= 0; i < len(alphabet); i++ {
			if currentKeyword == alphabet[i] {
				currentKeywordIndex = i
				break
			}
		}

		decipherIndex := currentIndex - currentKeywordIndex
		if decipherIndex < 0 {
			decipherIndex += len(alphabet)
		}
		// decipherIndex = decipherIndex % len(alphabet)
		
		// fmt.Printf("%c-%c %-2v - %-2v %c\n", currentKeyword, current, currentKeywordIndex, currentIndex, alphabet[decipherIndex])
		fmt.Printf("%c",alphabet[decipherIndex])
	}
	fmt.Println()
}