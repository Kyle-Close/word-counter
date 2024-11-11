package main

import "fmt"

func PrintWordCountList(wordCountList []WordCount) {
	for _, v := range wordCountList {
		fmt.Printf("%v: %v\n", v.word, v.count)
	}
}
