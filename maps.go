package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	mapOfWords := make(map[string]int)
	for _, word := range strings.Fields(s) {
		mapOfWords[word]++
	}
	return mapOfWords
}

func main() {
	wc.Test(wordCount)
}
