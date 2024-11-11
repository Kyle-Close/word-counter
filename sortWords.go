package main

import "sort"

type WordCount struct {
	word  string
	count int
}

// Returns a list of WordCount structs sorted by count
func SortWords(wordMap map[string]int) []WordCount {
	result := make([]WordCount, 0, len(wordMap))

	for i, v := range wordMap {
		result = append(result, WordCount{word: i, count: v})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].count == result[j].count {
			// result will be based off name in this case
			return result[i].word < result[j].word
		}
		return result[i].count > result[j].count
	})

	return result
}
