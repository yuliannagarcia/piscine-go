package piscine

import (
	"fmt"
	"sort"
	"strings"
)

func MaxWordCountN(s string, n int) map[string]int {
	// Split the input string into words using spaces as the delimiter
	words := strings.Fields(s)

	// Create a slice to store the lengths of words
	wordLengths := make([]int, len(words))

	// Calculate and store the lengths of each word
	for i, word := range words {
		wordLengths[i] = len(word)
	}

	// Sort the word lengths in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(wordLengths)))

	// Create a map to store the top N word lengths by count
	topN := make(map[string]int)

	// Iterate over the sorted word lengths to find the top N lengths
	for i := 0; i < n && i < len(wordLengths); i++ {
		maxLength := wordLengths[i]

		// Find all words with the maxLength
		for _, word := range words {
			if len(word) == maxLength {
				topN[word]++
			}
		}
	}

	return topN
}

func piscine() {
	input := "xx   x xxxxx xx x xxxx     xx   xxxx x xx xx x xxx x x xx  x xxx xxxx  xxxxx x    xxx  xxx       x  "
	result := MaxWordCountN(input, 1)
	fmt.Println(result)
}
