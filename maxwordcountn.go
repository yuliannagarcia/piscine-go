package piscine

import (
	"sort"
	"strings"
)

// MaxWordCountN returns a map of the n words that occur the most in a string text.
func MaxWordCountN(text string, n int) map[string]int {
	// Split the text into words using strings.Fields
	words := strings.Fields(text)

	// Create a map to store word counts
	wordCounts := make(map[string]int)

	// Count word occurrences
	for _, word := range words {
		// Remove leading/trailing punctuation and trim spaces
		word = strings.Trim(word, ".,!?\n\t\r\v\f\"'()[]{}")
		// Increment the word count
		wordCounts[word]++
	}

	// Create a list of word-count pairs for sorting
	wordList := make([]struct {
		word  string
		count int
	}, len(wordCounts))

	i := 0
	for word, count := range wordCounts {
		wordList[i] = struct {
			word  string
			count int
		}{word, count}
		i++
	}

	// Sort the list by word count (descending) and word (ascending)
	sortWordList(wordList)

	// Create a result map with the top n words and their counts
	result := make(map[string]int)
	for i := 0; i < n && i < len(wordList); i++ {
		result[wordList[i].word] = wordList[i].count
	}

	return result
}

// sortWordList sorts a list of word-count pairs by count (descending) and word (ascending).
func sortWordList(list []struct {
	word  string
	count int
},
) {
	// Sorting function
	sortFunc := func(i, j int) bool {
		// Sort by count (descending)
		if list[i].count != list[j].count {
			return list[i].count > list[j].count
		}
		// If counts are equal, sort by word (ascending)
		return list[i].word < list[j].word
	}

	// Sort the list using the sorting function
	sort.SliceStable(list, sortFunc)
}
