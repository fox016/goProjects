package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordToCount := make(map[string]int)
	for _, word := range words {
		_, exists := wordToCount[word]
		if exists {
			wordToCount[word]++
		} else {
			wordToCount[word]=1
		}
	}
	return wordToCount
}

func main() {
	m := WordCount("I like the fact that the pickle that is over there is not the same pickle as the one that is in that jar")
	fmt.Println(m)
}
