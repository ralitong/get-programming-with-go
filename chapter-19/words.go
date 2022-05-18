package main

import (
	"fmt"
	"sort"
	"strings"
)

func tokenize(paragraph string) []string {
	paragraphWithoutNewlines := strings.ReplaceAll(paragraph, "\n", " ")
	lowerCaseParagraph := strings.ToLower(paragraphWithoutNewlines)
	return strings.Fields(lowerCaseParagraph)
}

func createFrequencyMap(tokens []string) map[string]int {
	sort.Strings(tokens)
	frequencies := make(map[string]int)
	for _, token := range tokens {
		frequencies[token]++
	}

	return frequencies
}

func displayFrequencyMap(frequencyMap map[string]int) {
	for word, frequency := range frequencyMap {
		fmt.Printf("%-14s: %-16d\n", word, frequency)
	}
}


func main() {
	// create a token slice
	// create a frequency map -- map[string]int
	// display the word count

	paragraph := `
	As far as eye could reach he saw nothing but the stems of the great plants about him
	receding in the violet shade, and far overhead the multiple transparency of huge leaves
	filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever
	he felt able he ran again; the ground continued soft and springy, covered with the same
	resilient weed which was the first thing his hands had touched in Malacandra. Once or
	twice a small red creature scuttled across his path, but otherwise there seemed to be no
	life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned
	and alone in a forest of unknown vegetation thousands or millions of miles beyond the
	reach or knowledge of man.
	`

	tokens := tokenize(paragraph)
	frequencies := createFrequencyMap(tokens)
	displayFrequencyMap(frequencies)

}