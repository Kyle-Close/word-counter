package main

import (
	"fmt"
	"os"
)

func main() {
	lim := HandleArgs()

	file, err := os.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fileContents := string(file)
	wordMap := CountWords(fileContents)

	sorted := SortWords(wordMap)

	// Extract top 10 words.
	if len(sorted) > lim {
		sorted = sorted[:lim]
	}

	PrintWordCountList(sorted)
}
